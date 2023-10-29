package validation

import (
	"time"

	"github.com/tiendc/go-validator/base"
	timeFunc "github.com/tiendc/go-validator/base/time"
)

const (
	timeType = "time"
)

func TimeEQ[T base.Time](v T, val time.Time) SingleValidator {
	return call2[T]("eq", timeType, "TargetValue", timeFunc.EQ[T])(v, val)
}

func TimeGT[T base.Time](v T, min time.Time) SingleValidator {
	return call2[T]("gt", timeType, "Min", timeFunc.GT[T])(v, min)
}

func TimeGTE[T base.Time](v T, min time.Time) SingleValidator {
	return call2[T]("gte", timeType, "Min", timeFunc.GTE[T])(v, min)
}

func TimeLT[T base.Time](v T, max time.Time) SingleValidator {
	return call2[T]("lt", timeType, "Max", timeFunc.LT[T])(v, max)
}

func TimeLTE[T base.Time](v T, max time.Time) SingleValidator {
	return call2[T]("lte", timeType, "Max", timeFunc.LTE[T])(v, max)
}

func TimeValid[T base.Time](v T) SingleValidator {
	return call1[T]("valid", timeType, timeFunc.Valid[T])(v)
}

func TimeRange[T base.Time](v T, min, max time.Time) SingleValidator {
	return call3[T]("range", timeType, "Min", "Max", timeFunc.Range[T])(v, min, max)
}

func TimeIn[T base.Time](v T, s ...time.Time) SingleValidator {
	return call2N[T]("in", timeType, "TargetValue", timeFunc.In[T])(v, s...)
}

func TimeNotIn[T base.Time](v T, s ...time.Time) SingleValidator {
	return call2N[T]("not_in", timeType, "TargetValue", timeFunc.NotIn[T])(v, s...)
}
