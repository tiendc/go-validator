package validation

import (
	"context"

	"github.com/tiendc/go-validator/base"
)

// call1 generics function to execute a validation function with 1 param
func call1[P any](
	errType, valueType string,
	validateFunc func(P) (bool, []base.ErrorParam),
) func(v P) SingleValidator {
	return func(v P) SingleValidator {
		return NewSingleValidator(func(ctx context.Context) Error {
			success, params := validateFunc(v)
			if success {
				return nil
			}
			return errorBuild(errType, valueType, v, nil, params...)
		})
	}
}

// call2 generics function to execute a validation function with 2 params
func call2[P1 any, P2 any](
	errType, valueType string, p2Name string,
	validateFunc func(P1, P2) (bool, []base.ErrorParam),
) func(P1, P2) SingleValidator {
	return func(v P1, p2 P2) SingleValidator {
		return NewSingleValidator(func(ctx context.Context) Error {
			success, params := validateFunc(v, p2)
			if success {
				return nil
			}
			params = append(params, base.ErrorParam{Key: p2Name, Value: p2})
			return errorBuild(errType, valueType, v, nil, params...)
		})
	}
}

// call2N generics function to execute a validation function with 1 fixed param and a variable param
func call2N[P1 any, P2 any](
	errType, valueType string, p2Name string,
	validateFunc func(P1, ...P2) (bool, []base.ErrorParam),
) func(P1, ...P2) SingleValidator {
	return func(v P1, p2s ...P2) SingleValidator {
		return NewSingleValidator(func(ctx context.Context) Error {
			success, params := validateFunc(v, p2s...)
			if success {
				return nil
			}
			params = append(params, base.ErrorParam{Key: p2Name, Value: p2s})
			return errorBuild(errType, valueType, v, nil, params...)
		})
	}
}

// call3 generics function to execute a validation function with 3 params
func call3[P1 any, P2 any, P3 any](
	errType, valueType string, p2Name, p3Name string,
	validateFunc func(P1, P2, P3) (bool, []base.ErrorParam),
) func(P1, P2, P3) SingleValidator {
	return func(v P1, p2 P2, p3 P3) SingleValidator {
		return NewSingleValidator(func(ctx context.Context) Error {
			success, params := validateFunc(v, p2, p3)
			if success {
				return nil
			}
			params = append(params, base.ErrorParam{Key: p2Name, Value: p2},
				base.ErrorParam{Key: p3Name, Value: p3})
			return errorBuild(errType, valueType, v, nil, params...)
		})
	}
}

// ptrCall1 generics function to execute a validation function with 1 pointer param
func ptrCall1[P any](
	errType, valueType string,
	validateFunc func(P) (bool, []base.ErrorParam),
) func(v *P) SingleValidator {
	return func(v *P) SingleValidator {
		return NewSingleValidator(func(ctx context.Context) Error {
			if v == nil {
				return nil
			}
			success, params := validateFunc(*v)
			if success {
				return nil
			}
			return errorBuild(errType, valueType, *v, nil, params...)
		})
	}
}

// ptrCall2 generics function to execute a validation function with 2 pointer params
func ptrCall2[P1 any, P2 any](
	errType, valueType string, p2Name string,
	validateFunc func(P1, P2) (bool, []base.ErrorParam),
) func(*P1, P2) SingleValidator {
	return func(v *P1, p2 P2) SingleValidator {
		return NewSingleValidator(func(ctx context.Context) Error {
			if v == nil {
				return nil
			}
			success, params := validateFunc(*v, p2)
			if success {
				return nil
			}
			params = append(params, base.ErrorParam{Key: p2Name, Value: p2})
			return errorBuild(errType, valueType, *v, nil, params...)
		})
	}
}

// ptrCall2N generics function to execute a validation function with 1 fixed param and a variable param
func ptrCall2N[P1 any, P2 any](
	errType, valueType string, p2Name string,
	validateFunc func(P1, ...P2) (bool, []base.ErrorParam),
) func(*P1, ...P2) SingleValidator {
	return func(v *P1, p2s ...P2) SingleValidator {
		return NewSingleValidator(func(ctx context.Context) Error {
			if v == nil {
				return nil
			}
			success, params := validateFunc(*v, p2s...)
			if success {
				return nil
			}
			params = append(params, base.ErrorParam{Key: p2Name, Value: p2s})
			return errorBuild(errType, valueType, *v, nil, params...)
		})
	}
}

// ptrCall3 generics function to execute a validation function with 3 pointer params
func ptrCall3[P1 any, P2 any, P3 any](
	errType, valueType string, p2Name, p3Name string,
	validateFunc func(P1, P2, P3) (bool, []base.ErrorParam),
) func(*P1, P2, P3) SingleValidator {
	return func(v *P1, p2 P2, p3 P3) SingleValidator {
		return NewSingleValidator(func(ctx context.Context) Error {
			if v == nil {
				return nil
			}
			success, params := validateFunc(*v, p2, p3)
			if success {
				return nil
			}
			params = append(params, base.ErrorParam{Key: p2Name, Value: p2},
				base.ErrorParam{Key: p3Name, Value: p3})
			return errorBuild(errType, valueType, *v, nil, params...)
		})
	}
}
