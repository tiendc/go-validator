package timevalidation

import (
	"time"

	"github.com/tiendc/go-validator/base"
)

func EQ[T base.Time](v T, val time.Time) (bool, []base.ErrorParam) {
	return v.Compare(val) == 0, nil
}

func GT[T base.Time](v T, min time.Time) (bool, []base.ErrorParam) {
	return v.Compare(min) > 0, nil
}

func GTE[T base.Time](v T, min time.Time) (bool, []base.ErrorParam) {
	return v.Compare(min) >= 0, nil
}

func LT[T base.Time](v T, min time.Time) (bool, []base.ErrorParam) {
	return v.Compare(min) < 0, nil
}

func LTE[T base.Time](v T, min time.Time) (bool, []base.ErrorParam) {
	return v.Compare(min) <= 0, nil
}

func Valid[T base.Time](v T) (bool, []base.ErrorParam) {
	return !v.IsZero(), nil
}

func Range[T base.Time](v T, min, max time.Time) (bool, []base.ErrorParam) {
	return v.Compare(min) >= 0 && v.Compare(max) <= 0, nil
}

func In[T base.Time](v T, s ...time.Time) (bool, []base.ErrorParam) {
	for i := range s {
		if v.Compare(s[i]) == 0 {
			return true, nil
		}
	}
	return false, nil
}

func NotIn[T base.Time](v T, s ...time.Time) (bool, []base.ErrorParam) {
	for i := range s {
		if v.Compare(s[i]) == 0 {
			return false, nil
		}
	}
	return true, nil
}
