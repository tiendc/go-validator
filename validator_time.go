package validation

import (
	"time"

	"github.com/tiendc/go-validator/base"
	timeFunc "github.com/tiendc/go-validator/base/time"
)

const (
	timeType = "time"
)

// TimeEQ validates the input time must equal to the specified value
func TimeEQ[T base.Time](v T, val time.Time) SingleValidator {
	return call2[T]("eq", timeType, "TargetValue", timeFunc.EQ[T])(v, val)
}

// TimeGT validates the input time must be greater than the specified value
func TimeGT[T base.Time](v T, min time.Time) SingleValidator {
	return call2[T]("gt", timeType, "Min", timeFunc.GT[T])(v, min)
}

// TimeGTE validates the input time must be greater than or equal to the specified value
func TimeGTE[T base.Time](v T, min time.Time) SingleValidator {
	return call2[T]("gte", timeType, "Min", timeFunc.GTE[T])(v, min)
}

// TimeLT validates the input time must be less than the specified value
func TimeLT[T base.Time](v T, max time.Time) SingleValidator {
	return call2[T]("lt", timeType, "Max", timeFunc.LT[T])(v, max)
}

// TimeLTE validates the input time must be less than or equal to the specified value
func TimeLTE[T base.Time](v T, max time.Time) SingleValidator {
	return call2[T]("lte", timeType, "Max", timeFunc.LTE[T])(v, max)
}

// TimeValid validates the input time must be not zero value
func TimeValid[T base.Time](v T) SingleValidator {
	return call1[T]("valid", timeType, timeFunc.Valid[T])(v)
}

// TimeRange validates the input time must be in the specified range
func TimeRange[T base.Time](v T, min, max time.Time) SingleValidator {
	return call3[T]("range", timeType, "Min", "Max", timeFunc.Range[T])(v, min, max)
}

// TimeIn validates the input time must be in the specified values
func TimeIn[T base.Time](v T, s ...time.Time) SingleValidator {
	return call2N[T]("in", timeType, "TargetValue", timeFunc.In[T])(v, s...)
}

// TimeNotIn validates the input time must be not in the specified values
func TimeNotIn[T base.Time](v T, s ...time.Time) SingleValidator {
	return call2N[T]("not_in", timeType, "TargetValue", timeFunc.NotIn[T])(v, s...)
}
