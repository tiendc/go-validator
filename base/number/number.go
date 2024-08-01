package numbervalidation

import (
	"github.com/tiendc/go-validator/base"
)

const (
	JsMaxSafeInt = 9007199254740991
	JsMinSafeInt = -9007199254740991
)

// EQ checks number must equal to another value
func EQ[T base.Number](v T, val T) (bool, []base.ErrorParam) {
	return v == val, nil
}

// GT checks number must be greater than another value
func GT[T base.Number](v T, min T) (bool, []base.ErrorParam) {
	return v > min, nil
}

// GTE checks number must be greater than or equal to another value
func GTE[T base.Number](v T, min T) (bool, []base.ErrorParam) {
	return v >= min, nil
}

// LT checks number must be less than another value
func LT[T base.Number](v T, max T) (bool, []base.ErrorParam) {
	return v < max, nil
}

// LTE checks number must be less than or equal to another value
func LTE[T base.Number](v T, max T) (bool, []base.ErrorParam) {
	return v <= max, nil
}

// Range checks number must be in a range
func Range[T base.Number](v T, min, max T) (bool, []base.ErrorParam) {
	return min <= v && v <= max, nil
}

// In checks number must be in a list
func In[T base.Number](v T, s ...T) (bool, []base.ErrorParam) {
	for i := range s {
		if v == s[i] {
			return true, nil
		}
	}
	return false, nil
}

// NotIn checks number must be not in a list
func NotIn[T base.Number](v T, s ...T) (bool, []base.ErrorParam) {
	for i := range s {
		if v == s[i] {
			return false, nil
		}
	}
	return true, nil
}

// DivisibleBy checks number must be divisible by a value
func DivisibleBy[T base.Int | base.UInt](v T, div T) (bool, []base.ErrorParam) {
	if div == 0 {
		return false, nil
	}
	return v%div == 0, nil
}

// JsSafeInt checks number must be a Javascript safe integer (max 2^53-1)
func JsSafeInt[T base.Int | base.UInt](v T) (bool, []base.ErrorParam) {
	vv := int64(v)
	return JsMinSafeInt <= vv && vv <= JsMaxSafeInt, nil
}
