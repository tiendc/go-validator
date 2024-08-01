package validation

import (
	"github.com/tiendc/go-validator/base"
	numFunc "github.com/tiendc/go-validator/base/number"
)

const (
	numType = "number"
)

// NumEQ validates the input number must equal to a value
func NumEQ[T base.Number](v *T, val T) SingleValidator {
	return ptrCall2[T]("eq", numType, "TargetValue", numFunc.EQ[T])(v, val)
}

// NumGT validates the input number must be greater than a value
func NumGT[T base.Number](v *T, min T) SingleValidator {
	return ptrCall2[T]("gt", numType, "Min", numFunc.GT[T])(v, min)
}

// NumGTE validates the input number must be greater than or equal to a value
func NumGTE[T base.Number](v *T, min T) SingleValidator {
	return ptrCall2[T]("gte", numType, "Min", numFunc.GTE[T])(v, min)
}

// NumLT validates the input number must be less than a value
func NumLT[T base.Number](v *T, max T) SingleValidator {
	return ptrCall2[T]("lt", numType, "Max", numFunc.LT[T])(v, max)
}

// NumLTE validates the input number must be less than or equal to a value
func NumLTE[T base.Number](v *T, max T) SingleValidator {
	return ptrCall2[T]("lte", numType, "Max", numFunc.LTE[T])(v, max)
}

// NumRange validates the input number must be in the specified range
func NumRange[T base.Number](v *T, min, max T) SingleValidator {
	return ptrCall3[T]("range", numType, "Min", "Max", numFunc.Range[T])(v, min, max)
}

// NumIn validates the input number must be in the specified values
func NumIn[T base.Number](v *T, s ...T) SingleValidator {
	return ptrCall2N[T]("in", numType, "TargetValue", numFunc.In[T])(v, s...)
}

// NumNotIn validates the input number must be not in the specified values
func NumNotIn[T base.Number](v *T, s ...T) SingleValidator {
	return ptrCall2N[T]("not_in", numType, "TargetValue", numFunc.NotIn[T])(v, s...)
}

// NumDivisibleBy validates the input number must be divisible by the specified value
func NumDivisibleBy[T base.Int | base.UInt](v *T, div T) SingleValidator {
	return ptrCall2[T]("divisible_by", numType, "TargetValue", numFunc.DivisibleBy[T])(v, div)
}

// NumJsSafeInt validates the input number must be a Javascript safe integer (max 2^53-1)
func NumJsSafeInt[T base.Int | base.UInt](v *T) SingleValidator {
	return ptrCall1[T]("js_safe_int", numType, numFunc.JsSafeInt[T])(v)
}
