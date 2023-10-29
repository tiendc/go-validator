package validation

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// nolint: unparam
func parse(s, layout string) time.Time {
	t, e := time.Parse(layout, s)
	if e != nil {
		panic(e)
	}
	return t
}

func Test_EQ(t *testing.T) {
	l1 := time.DateTime
	errs := TimeEQ(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1)).Exec()
	assert.Equal(t, 0, len(errs))

	errs = TimeEQ(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:11", l1)).Exec()
	assert.Equal(t, "eq", errs[0].Type())
}

func Test_TimeGT_TimeGTE(t *testing.T) {
	l1 := time.DateTime
	errs := TimeGT(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:09", l1)).Exec()
	assert.Equal(t, 0, len(errs))
	errs = TimeGTE(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1)).Exec()
	assert.Equal(t, 0, len(errs))

	errs = TimeGT(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1)).Exec()
	assert.Equal(t, "gt", errs[0].Type())
	errs = TimeGTE(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:11", l1)).Exec()
	assert.Equal(t, "gte", errs[0].Type())
}

func Test_TimeLT_TimeLTE(t *testing.T) {
	l1 := time.DateTime
	errs := TimeLT(parse("2023-10-01 10:10:09", l1), parse("2023-10-01 10:10:10", l1)).Exec()
	assert.Equal(t, 0, len(errs))
	errs = TimeLTE(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1)).Exec()
	assert.Equal(t, 0, len(errs))

	errs = TimeLT(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1)).Exec()
	assert.Equal(t, "lt", errs[0].Type())
	errs = TimeLTE(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:09", l1)).Exec()
	assert.Equal(t, "lte", errs[0].Type())
}

func Test_TimeValid(t *testing.T) {
	l1 := time.DateTime
	errs := TimeValid(parse("2023-10-01 10:10:10", l1)).Exec()
	assert.Equal(t, 0, len(errs))

	errs = TimeValid(time.Time{}).Exec()
	assert.Equal(t, "valid", errs[0].Type())
}

func Test_TimeRange(t *testing.T) {
	l1 := time.DateTime
	errs := TimeRange(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:11", l1)).Exec()
	assert.Equal(t, 0, len(errs))

	errs = TimeRange(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:11", l1), parse("2023-10-01 10:10:12", l1)).Exec()
	assert.Equal(t, "range", errs[0].Type())
}

func Test_TimeIn(t *testing.T) {
	l1 := time.DateTime
	errs := TimeIn(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:11", l1)).Exec()
	assert.Equal(t, 0, len(errs))

	errs = TimeIn(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:11", l1), parse("2023-10-01 10:10:12", l1)).Exec()
	assert.Equal(t, "in", errs[0].Type())
}

func Test_TimeNotIn(t *testing.T) {
	l1 := time.DateTime
	errs := TimeNotIn(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:11", l1), parse("2023-10-01 10:10:12", l1)).Exec()
	assert.Equal(t, 0, len(errs))

	errs = TimeNotIn(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:11", l1), parse("2023-10-01 10:10:10", l1)).Exec()
	assert.Equal(t, "not_in", errs[0].Type())
}
