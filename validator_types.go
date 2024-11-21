package validation

import (
	"context"
	"fmt"
)

// Validator interface represents a validator object
type Validator interface {
	Validate(context.Context) Errors
	OnError(...ErrorMod) Validator
}

// baseValidator base validator
type baseValidator struct {
	errMods []ErrorMod
}

// applyErrMods applies mods on the target error
func (b *baseValidator) applyErrMods(err Error) {
	for _, mod := range b.errMods {
		mod(err)
	}
}

// applyErrModsWithGrouping applies mods on the target error.
// If the input has more than 1 error, this creates a `group` error and applies mods on it.
func (b *baseValidator) applyErrModsWithGrouping(errs Errors) Errors {
	if len(b.errMods) == 0 || len(errs) == 0 {
		return errs
	}
	if len(errs) > 1 {
		errs = []Error{errorBuild("group", "", nil, errs)}
	}
	b.applyErrMods(errs[0])
	return errs
}

// SingleValidator interface represents a validator that performs a single validation
type SingleValidator interface {
	Validator
}

// NewSingleValidator creates a new SingleValidator
func NewSingleValidator(execFn func(ctx context.Context) Error) SingleValidator {
	return &singleValidator{
		execFn: execFn,
	}
}

// singleValidator implementation of SingleValidator interface
type singleValidator struct {
	baseValidator
	execFn func(ctx context.Context) Error
}

// OnError implementation of Validator interface
func (v *singleValidator) OnError(mods ...ErrorMod) Validator {
	v.errMods = mods
	return v
}

// Validate executes the validator
func (v *singleValidator) Validate(ctx context.Context) Errors {
	err := v.execFn(ctx)
	if err == nil {
		return nil
	}
	v.applyErrMods(err)
	return []Error{err}
}

// CondValidator interface represents a validator that performs multiple validations based on
// specified conditions.
type CondValidator interface {
	Validator
	ValidateWithCond(context.Context) (bool, Errors)
}

// SingleCondValidator validator that accepts only one condition
type SingleCondValidator interface {
	CondValidator
	Then(validators ...Validator) SingleCondValidator
	Else(validators ...Validator) SingleCondValidator
}

// singleCondValidator implementation of SingleCondValidator
type singleCondValidator struct {
	baseValidator
	conditions     []any
	thenValidators []Validator
	elseValidators []Validator
}

// NewSingleCondValidator creates a new SingleCondValidator
func NewSingleCondValidator(conditions ...any) SingleCondValidator {
	return &singleCondValidator{conditions: conditions}
}

func (c *singleCondValidator) Then(validators ...Validator) SingleCondValidator {
	c.thenValidators = validators
	return c
}

func (c *singleCondValidator) Else(validators ...Validator) SingleCondValidator {
	c.elseValidators = validators
	return c
}

func (c *singleCondValidator) OnError(errMods ...ErrorMod) Validator {
	c.errMods = errMods
	return c
}

func (c *singleCondValidator) Validate(ctx context.Context) Errors {
	_, errs := c.ValidateWithCond(ctx)
	return c.applyErrModsWithGrouping(errs)
}

func (c *singleCondValidator) ValidateWithCond(ctx context.Context) (bool, Errors) {
	validators := c.thenValidators
	match := c.match(ctx)
	if !match {
		validators = c.elseValidators
	}
	if len(validators) == 0 {
		return match, nil
	}
	return match, execValidators(ctx, validators, false)
}

func (c *singleCondValidator) match(ctx context.Context) bool {
	if len(c.conditions) == 0 {
		return false
	}
	for _, cond := range c.conditions {
		boolVal, ok := cond.(bool)
		if ok {
			if !boolVal {
				return false
			}
			continue
		}
		validator, ok := cond.(Validator)
		if ok {
			errs := validator.Validate(ctx)
			if len(errs) > 0 {
				return false
			}
			continue
		}
		panic(fmt.Errorf("%w: only 'bool' or 'validator' allowed", ErrTypeUnsupported))
	}
	return true
}

// MultiCondValidator validator that accepts multiple conditions
type MultiCondValidator interface {
	CondValidator
	Default(validators ...Validator) MultiCondValidator
}

// multiCondValidator implementation of MultiCondValidator
type multiCondValidator struct {
	baseValidator
	conditions        []SingleCondValidator
	defaultValidators []Validator
}

// NewMultiCondValidator creates a new MultiCondValidator
func NewMultiCondValidator(conditions ...SingleCondValidator) MultiCondValidator {
	return &multiCondValidator{conditions: conditions}
}

func (c *multiCondValidator) Default(validators ...Validator) MultiCondValidator {
	c.defaultValidators = validators
	return c
}

func (c *multiCondValidator) OnError(mods ...ErrorMod) Validator {
	c.errMods = mods
	return c
}

func (c *multiCondValidator) Validate(ctx context.Context) Errors {
	_, errs := c.ValidateWithCond(ctx)
	return c.applyErrModsWithGrouping(errs)
}

func (c *multiCondValidator) ValidateWithCond(ctx context.Context) (bool, Errors) {
	for _, v := range c.conditions {
		match, errs := v.ValidateWithCond(ctx)
		if match {
			return true, errs
		}
	}
	// No match condition, executes the default ones
	return false, execValidators(ctx, c.defaultValidators, false)
}

// ItemValidator validator to collect input validators via its Validate() func
type ItemValidator interface {
	Validate(validators ...Validator)
	Group(validators ...Validator) SingleValidator
	OneOf(validators ...Validator) SingleValidator
	ExactOneOf(validators ...Validator) SingleValidator
	NotOf(validators ...Validator) SingleValidator
}

// itemValidator implements ItemValidator interface
type itemValidator struct {
	ItemValidator
	validators []Validator
}

func (iv *itemValidator) Validate(validators ...Validator) {
	iv.validators = append(iv.validators, validators...)
}

func (iv *itemValidator) Group(validators ...Validator) SingleValidator {
	group := Group(validators...)
	iv.validators = append(iv.validators, group)
	return group
}

func (iv *itemValidator) OneOf(validators ...Validator) SingleValidator {
	oneOf := OneOf(validators...)
	iv.validators = append(iv.validators, oneOf)
	return oneOf
}

func (iv *itemValidator) ExactOneOf(validators ...Validator) SingleValidator {
	exactOneOf := ExactOneOf(validators...)
	iv.validators = append(iv.validators, exactOneOf)
	return exactOneOf
}

func (iv *itemValidator) NotOf(validators ...Validator) SingleValidator {
	notOf := NotOf(validators...)
	iv.validators = append(iv.validators, notOf)
	return notOf
}

func (iv *itemValidator) get() []Validator {
	return iv.validators
}

// SliceContentValidator validator that validates slice elements
type SliceContentValidator[T any, S ~[]T] interface {
	Validator
	ForEach(fn func(element T, index int, elemValidator ItemValidator)) SliceContentValidator[T, S]
}

// sliceContentValidator implementation of SliceContentValidator
type sliceContentValidator[T any, S ~[]T] struct {
	baseValidator
	slice             S
	elemValidatorFunc func(T, int, ItemValidator)
}

// NewSliceElemValidator creates a new SliceContentValidator
func NewSliceElemValidator[T any, S ~[]T](slice S) SliceContentValidator[T, S] {
	return &sliceContentValidator[T, S]{slice: slice}
}

func (c *sliceContentValidator[T, S]) ForEach(fn func(T, int, ItemValidator)) SliceContentValidator[T, S] {
	c.elemValidatorFunc = fn
	return c
}

func (c *sliceContentValidator[T, S]) OnError(errMods ...ErrorMod) Validator {
	c.errMods = errMods
	return c
}

func (c *sliceContentValidator[T, S]) Validate(ctx context.Context) Errors {
	if len(c.slice) == 0 {
		return nil
	}
	elemValidator := &itemValidator{}
	for i, elem := range c.slice {
		c.elemValidatorFunc(elem, i, elemValidator)
	}
	errs := execValidators(ctx, elemValidator.get(), false)
	return c.applyErrModsWithGrouping(errs)
}
