package validation

import "fmt"

// Validator interface represents a validator object
type Validator interface {
	Exec() Errors
}

// SingleValidator interface represents a validator that performs a single validation
type SingleValidator interface {
	Validator
	OnError(...ErrorMod) SingleValidator
}

// NewSingleValidator creates a new SingleValidator
func NewSingleValidator(execFn func() Error) SingleValidator {
	return &singleValidator{
		execFn: execFn,
	}
}

// singleValidator implementation of SingleValidator interface
type singleValidator struct {
	execFn  func() Error
	errMods []ErrorMod
}

// Exec executes the validator
func (v *singleValidator) Exec() Errors {
	err := v.execFn()
	if err == nil {
		return nil
	}
	for _, fn := range v.errMods {
		fn(err)
	}
	return []Error{err}
}

// OnError sets modifiers for validation error when it comes
func (v *singleValidator) OnError(mods ...ErrorMod) SingleValidator {
	v.errMods = append(v.errMods, mods...)
	return v
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
	conditions     []any
	thenValidators []Validator
	elseValidators []Validator
}

// NewSingleCondValidator creates a new SingleCondValidator
func NewSingleCondValidator(conditions ...any) SingleCondValidator {
	return &singleCondValidator{conditions: conditions}
}

func (c *singleCondValidator) Then(validators ...Validator) SingleCondValidator {
	c.thenValidators = append(c.thenValidators, validators...)
	return c
}

func (c *singleCondValidator) Else(validators ...Validator) SingleCondValidator {
	c.elseValidators = append(c.elseValidators, validators...)
	return c
}

func (c *singleCondValidator) Exec() Errors {
	_, errs := c.ExecEx()
	return errs
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
	conditions        []SingleCondValidator
	defaultValidators []Validator
}

// NewMultiCondValidator creates a new MultiCondValidator
func NewMultiCondValidator(conditions ...SingleCondValidator) MultiCondValidator {
	return &multiCondValidator{conditions: conditions}
}

func (c *multiCondValidator) Default(validators ...Validator) MultiCondValidator {
	c.defaultValidators = append(c.defaultValidators, validators...)
	return c
}

func (c *multiCondValidator) Exec() Errors {
	_, errs := c.ExecEx()
	return errs
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
