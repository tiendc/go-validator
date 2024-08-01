package validation

// Validate executes given validators to make result.
func Validate(validators ...Validator) Errors {
	return execValidators(validators, false)
}

// Group groups the given validators into one.
// In case there are errors, only one error will be returned.
func Group(validators ...Validator) SingleValidator {
	return NewSingleValidator(func() Error {
		if len(validators) == 0 {
			return nil
		}
		errs := execValidators(validators, false)
		if len(errs) == 0 {
			return nil
		}
		return errorBuild("group", "", nil, errs)
	})
}

// OneOf checks if the target value satisfies one of the given validators.
// When a validator passes, the remaining ones will be skipped.
func OneOf(validators ...Validator) SingleValidator {
	return NewSingleValidator(func() Error {
		if len(validators) == 0 {
			return nil
		}
		wrapErrs := Errors{}
		for _, v := range validators {
			errs := v.Exec()
			if len(errs) == 0 {
				return nil // return nil when one passes
			}
			wrapErrs = append(wrapErrs, errs...)
		}
		return errorBuild("one_of", "", nil, wrapErrs)
	})
}

// ExactOneOf checks if the target value satisfies only one of the given validators.
// This returns error when there is not one or more than one validator pass.
func ExactOneOf(validators ...Validator) SingleValidator {
	return NewSingleValidator(func() Error {
		numValidatorPass := 0
		wrapErrs := Errors{}
		for _, v := range validators {
			errs := v.Exec()
			if len(errs) == 0 {
				numValidatorPass++
			} else {
				wrapErrs = append(wrapErrs, errs...)
			}
			if numValidatorPass > 1 {
				break
			}
		}
		if numValidatorPass == 1 {
			return nil
		}
		return errorBuild("exact_one_of", "", nil, wrapErrs)
	})
}

// NotOf checks the target value not satisfy any of the given validators.
// When a validator passes, an error will be returned and the remaining checks will be skipped.
func NotOf(validators ...Validator) SingleValidator {
	return NewSingleValidator(func() Error {
		for _, v := range validators {
			errs := v.Exec()
			if len(errs) == 0 {
				return errorBuild("not_of", "", nil, nil)
			}
		}
		return nil
	})
}

// If a bare check validation convenient for validating custom data such as
// `If(myTime.Before(dueDate)).OnError(vld.SetCustomKey("MY_ERR_KEY"))`.
// Deprecated: use `Must` instead
func If(cond bool) SingleValidator {
	return NewSingleValidator(func() Error {
		if cond {
			return nil
		}
		return errorBuild("if", "", cond, nil)
	})
}

// Must a bare check validation convenient for validating custom data such as
// `Must(myTime.Before(dueDate)).OnError(vld.SetCustomKey("MY_ERR_KEY"))`.
func Must(cond bool) SingleValidator {
	return NewSingleValidator(func() Error {
		if cond {
			return nil
		}
		return errorBuild("must", "", cond, nil)
	})
}

// When works like a `if...then...else` statement
func When(conditions ...any) SingleCondValidator {
	return NewSingleCondValidator(conditions...)
}

// Case works like a `switch...case` statement
func Case(conditions ...SingleCondValidator) MultiCondValidator {
	return NewMultiCondValidator(conditions...)
}

// execValidators executes a list of validators and collect errors as result
// nolint: unparam
func execValidators(validators []Validator, stopOnError bool) Errors {
	errs := make(Errors, 0, len(validators))
	for _, v := range validators {
		hasErr := false
		for _, e := range v.Exec() {
			if e == nil {
				continue
			}
			hasErr = true
			errs = append(errs, e)
		}
		if stopOnError && hasErr {
			break
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}
