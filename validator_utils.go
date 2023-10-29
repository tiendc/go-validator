package validation

import (
	"github.com/tiendc/go-validator/base"
)

func call1[P any](
	errType, valueType string,
	validateFunc func(P) (bool, []base.ErrorParam),
) func(v P) SingleValidator {
	return func(v P) SingleValidator {
		return NewSingleValidator(func() Error {
			success, params := validateFunc(v)
			if success {
				return nil
			}
			return errorBuild(errType, valueType, v, nil, params...)
		})
	}
}

func call2[P1 any, P2 any](
	errType, valueType string, p2Name string,
	validateFunc func(P1, P2) (bool, []base.ErrorParam),
) func(P1, P2) SingleValidator {
	return func(v P1, p2 P2) SingleValidator {
		return NewSingleValidator(func() Error {
			success, params := validateFunc(v, p2)
			if success {
				return nil
			}
			params = append(params, base.ErrorParam{Key: p2Name, Value: p2})
			return errorBuild(errType, valueType, v, nil, params...)
		})
	}
}

func call2N[P1 any, P2 any](
	errType, valueType string, p2Name string,
	validateFunc func(P1, ...P2) (bool, []base.ErrorParam),
) func(P1, ...P2) SingleValidator {
	return func(v P1, p2s ...P2) SingleValidator {
		return NewSingleValidator(func() Error {
			success, params := validateFunc(v, p2s...)
			if success {
				return nil
			}
			params = append(params, base.ErrorParam{Key: p2Name, Value: p2s})
			return errorBuild(errType, valueType, v, nil, params...)
		})
	}
}

func call3[P1 any, P2 any, P3 any](
	errType, valueType string, p2Name, p3Name string,
	validateFunc func(P1, P2, P3) (bool, []base.ErrorParam),
) func(P1, P2, P3) SingleValidator {
	return func(v P1, p2 P2, p3 P3) SingleValidator {
		return NewSingleValidator(func() Error {
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

func ptrCall1[P any](
	errType, valueType string,
	validateFunc func(P) (bool, []base.ErrorParam),
) func(v *P) SingleValidator {
	return func(v *P) SingleValidator {
		return NewSingleValidator(func() Error {
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

func ptrCall2[P1 any, P2 any](
	errType, valueType string, p2Name string,
	validateFunc func(P1, P2) (bool, []base.ErrorParam),
) func(*P1, P2) SingleValidator {
	return func(v *P1, p2 P2) SingleValidator {
		return NewSingleValidator(func() Error {
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

func ptrCall2N[P1 any, P2 any](
	errType, valueType string, p2Name string,
	validateFunc func(P1, ...P2) (bool, []base.ErrorParam),
) func(*P1, ...P2) SingleValidator {
	return func(v *P1, p2s ...P2) SingleValidator {
		return NewSingleValidator(func() Error {
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

func ptrCall3[P1 any, P2 any, P3 any](
	errType, valueType string, p2Name, p3Name string,
	validateFunc func(P1, P2, P3) (bool, []base.ErrorParam),
) func(*P1, P2, P3) SingleValidator {
	return func(v *P1, p2 P2, p3 P3) SingleValidator {
		return NewSingleValidator(func() Error {
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
