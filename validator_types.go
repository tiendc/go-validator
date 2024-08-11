package validation

import "fmt"

// Validator interface represents a validator object
type Validator interface {
	Exec() Errors
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
func NewSingleValidator(execFn func() Error) SingleValidator {
	return &singleValidator{
		execFn: execFn,
	}
}

// singleValidator implementation of SingleValidator interface
type singleValidator struct {
	baseValidator
	execFn func() Error
}

// OnError implementation of Validator interface
func (v *singleValidator) OnError(mods ...ErrorMod) Validator {
	v.errMods = mods
	return v
}

// Exec executes the validator
func (v *singleValidator) Exec() Errors {
	err := v.execFn()
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
	ExecEx() (bool, Errors)
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

func (c *singleCondValidator) Exec() Errors {
	_, errs := c.ExecEx()
	return c.applyErrModsWithGrouping(errs)
}

func (c *singleCondValidator) ExecEx() (bool, Errors) {
	validators := c.thenValidators
	match := c.match()
	if !match {
		validators = c.elseValidators
	}
	if len(validators) == 0 {
		return match, nil
	}
	return match, execValidators(validators, false)
}

func (c *singleCondValidator) match() bool {
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
			errs := validator.Exec()
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

func (c *multiCondValidator) Exec() Errors {
	_, errs := c.ExecEx()
	return c.applyErrModsWithGrouping(errs)
}

func (c *multiCondValidator) ExecEx() (bool, Errors) {
	for _, v := range c.conditions {
		match, errs := v.ExecEx()
		if match {
			return true, errs
		}
	}
	// No match condition, executes the default ones
	return false, execValidators(c.defaultValidators, false)
}
