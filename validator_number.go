package validation

import (
	"github.com/tiendc/go-validator/base"
	numFunc "github.com/tiendc/go-validator/base/number"
)

const (
	numType = "number"
)

func NumEQ[T base.Number](v *T, val T) SingleValidator {
	return ptrCall2[T]("eq", numType, "TargetValue", numFunc.EQ[T])(v, val)
}

func NumGT[T base.Number](v *T, min T) SingleValidator {
	return ptrCall2[T]("gt", numType, "Min", numFunc.GT[T])(v, min)
}

func NumGTE[T base.Number](v *T, min T) SingleValidator {
	return ptrCall2[T]("gte", numType, "Min", numFunc.GTE[T])(v, min)
}

func NumLT[T base.Number](v *T, max T) SingleValidator {
	return ptrCall2[T]("lt", numType, "Max", numFunc.LT[T])(v, max)
}

func NumLTE[T base.Number](v *T, max T) SingleValidator {
	return ptrCall2[T]("lte", numType, "Max", numFunc.LTE[T])(v, max)
}

func NumRange[T base.Number](v *T, min, max T) SingleValidator {
	return ptrCall3[T]("range", numType, "Min", "Max", numFunc.Range[T])(v, min, max)
}

func NumIn[T base.Number](v *T, s ...T) SingleValidator {
	return ptrCall2N[T]("in", numType, "TargetValue", numFunc.In[T])(v, s...)
}

func NumNotIn[T base.Number](v *T, s ...T) SingleValidator {
	return ptrCall2N[T]("not_in", numType, "TargetValue", numFunc.NotIn[T])(v, s...)
}

func NumDivisibleBy[T base.Int | base.UInt](v *T, div T) SingleValidator {
	return ptrCall2[T]("divisible_by", numType, "TargetValue", numFunc.DivisibleBy[T])(v, div)
}

func NumJsSafeInt[T base.Int | base.UInt](v *T) SingleValidator {
	return ptrCall1[T]("js_safe_int", numType, numFunc.JsSafeInt[T])(v)
}
