package timevalidation

import (
	"time"

	"github.com/tiendc/go-validator/base"
)

// EQ checks a time must equal to a time value
func EQ[T base.Time](v T, val time.Time) (bool, []base.ErrorParam) {
	return v.Compare(val) == 0, nil
}

// GT checks a time must be greater than a time value
func GT[T base.Time](v T, min time.Time) (bool, []base.ErrorParam) {
	return v.Compare(min) > 0, nil
}

// GTE checks a time must be greater than or equal to a time value
func GTE[T base.Time](v T, min time.Time) (bool, []base.ErrorParam) {
	return v.Compare(min) >= 0, nil
}

// LT checks a time must be less than a time value
func LT[T base.Time](v T, min time.Time) (bool, []base.ErrorParam) {
	return v.Compare(min) < 0, nil
}

// LTE checks a time must be less than or equal to a time value
func LTE[T base.Time](v T, min time.Time) (bool, []base.ErrorParam) {
	return v.Compare(min) <= 0, nil
}

// Valid checks a time must be non-zero value
func Valid[T base.Time](v T) (bool, []base.ErrorParam) {
	return !v.IsZero(), nil
}

// Range checks a time must be in a range
func Range[T base.Time](v T, min, max time.Time) (bool, []base.ErrorParam) {
	return v.Compare(min) >= 0 && v.Compare(max) <= 0, nil
}

// In checks a time must be in a list
func In[T base.Time](v T, s ...time.Time) (bool, []base.ErrorParam) {
	for i := range s {
		if v.Compare(s[i]) == 0 {
			return true, nil
		}
	}
	return false, nil
}

// NotIn checks a time must be not in a list
func NotIn[T base.Time](v T, s ...time.Time) (bool, []base.ErrorParam) {
	for i := range s {
		if v.Compare(s[i]) == 0 {
			return false, nil
		}
	}
	return true, nil
}
