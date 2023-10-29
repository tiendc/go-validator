package numbervalidation

import (
	"github.com/tiendc/go-validator/base"
)

const (
	JsMaxSafeInt = 9007199254740991
	JsMinSafeInt = -9007199254740991
)

func EQ[T base.Number](v T, val T) (bool, []base.ErrorParam) {
	return v == val, nil
}

func GT[T base.Number](v T, min T) (bool, []base.ErrorParam) {
	return v > min, nil
}

func GTE[T base.Number](v T, min T) (bool, []base.ErrorParam) {
	return v >= min, nil
}

func LT[T base.Number](v T, max T) (bool, []base.ErrorParam) {
	return v < max, nil
}

func LTE[T base.Number](v T, max T) (bool, []base.ErrorParam) {
	return v <= max, nil
}

func Range[T base.Number](v T, min, max T) (bool, []base.ErrorParam) {
	return min <= v && v <= max, nil
}

func In[T base.Number](v T, s ...T) (bool, []base.ErrorParam) {
	for i := range s {
		if v == s[i] {
			return true, nil
		}
	}
	return false, nil
}

func NotIn[T base.Number](v T, s ...T) (bool, []base.ErrorParam) {
	for i := range s {
		if v == s[i] {
			return false, nil
		}
	}
	return true, nil
}

func DivisibleBy[T base.Int | base.UInt](v T, div T) (bool, []base.ErrorParam) {
	if div == 0 {
		return false, nil
	}
	return v%div == 0, nil
}

func JsSafeInt[T base.Int | base.UInt](v T) (bool, []base.ErrorParam) {
	vv := int64(v)
	return JsMinSafeInt <= vv && vv <= JsMaxSafeInt, nil
}
