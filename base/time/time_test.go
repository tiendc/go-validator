package timevalidation

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func parse(s, layout string) time.Time {
	t, e := time.Parse(layout, s)
	if e != nil {
		panic(e)
	}
	return t
}

func Test_EQ(t *testing.T) {
	l1 := time.DateTime
	l2 := time.DateOnly
	assert.True(t, gofn.Head(EQ(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1))))
	assert.True(t, gofn.Head(EQ(parse("2023-10-01", l2), parse("2023-10-01", l2))))

	assert.False(t, gofn.Head(EQ(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:11", l1))))
	assert.False(t, gofn.Head(EQ(parse("2023-10-01", l2), parse("2023-10-02", l2))))
}

func Test_GT_GTE(t *testing.T) {
	l1 := time.DateTime
	l2 := time.DateOnly
	assert.True(t, gofn.Head(GT(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:09", l1))))
	assert.True(t, gofn.Head(GT(parse("2023-10-02 10:10:10", l1), parse("2023-10-01 10:10:10", l1))))
	assert.True(t, gofn.Head(GT(parse("2023-10-02", l2), parse("2023-10-01", l2))))
	assert.True(t, gofn.Head(GTE(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1))))
	assert.True(t, gofn.Head(GTE(parse("2023-10-01", l2), parse("2023-10-01", l2))))

	assert.False(t, gofn.Head(GT(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1))))
	assert.False(t, gofn.Head(GT(parse("2023-10-01", l2), parse("2023-10-01", l2))))
	assert.False(t, gofn.Head(GTE(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:11", l1))))
	assert.False(t, gofn.Head(GTE(parse("2023-10-01", l2), parse("2023-10-02", l2))))
}

func Test_LT_LTE(t *testing.T) {
	l1 := time.DateTime
	l2 := time.DateOnly
	assert.True(t, gofn.Head(LT(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:11", l1))))
	assert.True(t, gofn.Head(LT(parse("2023-10-01 10:10:10", l1), parse("2023-10-02 10:10:10", l1))))
	assert.True(t, gofn.Head(LT(parse("2023-10-01", l2), parse("2023-10-02", l2))))
	assert.True(t, gofn.Head(LTE(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1))))
	assert.True(t, gofn.Head(LTE(parse("2023-10-01", l2), parse("2023-10-01", l2))))

	assert.False(t, gofn.Head(LT(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1))))
	assert.False(t, gofn.Head(LT(parse("2023-10-01", l2), parse("2023-10-01", l2))))
	assert.False(t, gofn.Head(LTE(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:09", l1))))
	assert.False(t, gofn.Head(LTE(parse("2023-10-02", l2), parse("2023-10-01", l2))))
}

func Test_Valid(t *testing.T) {
	l1 := time.DateTime
	l2 := time.DateOnly
	assert.True(t, gofn.Head(Valid(parse("2023-10-01 10:10:10", l1))))
	assert.True(t, gofn.Head(Valid(parse("2023-10-01", l2))))
	assert.False(t, gofn.Head(Valid(time.Time{})))
}

func Test_Range(t *testing.T) {
	l1 := time.DateTime
	l2 := time.DateOnly
	assert.True(t, gofn.Head(Range(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1))))
	assert.True(t, gofn.Head(Range(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:09", l1), parse("2023-10-01 10:10:11", l1))))
	assert.True(t, gofn.Head(Range(parse("2023-10-02 10:10:10", l1),
		parse("2023-10-01 10:10:10", l1), parse("2023-10-03 10:10:10", l1))))
	assert.True(t, gofn.Head(Range(parse("2023-10-01", l2),
		parse("2023-10-01", l2), parse("2023-10-01", l2))))
	assert.True(t, gofn.Head(Range(parse("2023-10-02", l2),
		parse("2023-10-01", l2), parse("2023-10-03", l2))))

	assert.False(t, gofn.Head(Range(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:11", l1), parse("2023-10-01 10:10:12", l1))))
	assert.False(t, gofn.Head(Range(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:08", l1), parse("2023-10-01 10:10:09", l1))))
	assert.False(t, gofn.Head(Range(parse("2023-10-01", l2),
		parse("2023-10-02", l2), parse("2023-10-03", l2))))
	assert.False(t, gofn.Head(Range(parse("2023-10-03", l2),
		parse("2023-10-01", l2), parse("2023-10-02", l2))))
}

func Test_In(t *testing.T) {
	l1 := time.DateTime
	l2 := time.DateOnly
	assert.True(t, gofn.Head(In(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1))))
	assert.True(t, gofn.Head(In(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:09", l1), parse("2023-10-01 10:10:10", l1))))
	assert.True(t, gofn.Head(In(parse("2023-10-02 10:10:10", l1),
		parse("2023-10-02 10:10:10", l1), parse("2023-10-03 10:10:10", l1))))
	assert.True(t, gofn.Head(In(parse("2023-10-01", l2), parse("2023-10-01", l2))))
	assert.True(t, gofn.Head(In(parse("2023-10-02", l2),
		parse("2023-10-01", l2), parse("2023-10-02", l2))))

	assert.False(t, gofn.Head(In(parse("2023-10-01 10:10:10", l1))))
	assert.False(t, gofn.Head(In(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:11", l1))))
	assert.False(t, gofn.Head(In(parse("2023-10-01", l2), parse("2023-10-02", l2))))
	assert.False(t, gofn.Head(In(parse("2023-10-01", l2),
		parse("2023-10-02", l2), parse("2023-10-03", l2))))
}

func Test_NotIn(t *testing.T) {
	l1 := time.DateTime
	l2 := time.DateOnly
	assert.True(t, gofn.Head(NotIn(parse("2023-10-01 10:10:10", l1))))
	assert.True(t, gofn.Head(NotIn(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:11", l1))))
	assert.True(t, gofn.Head(NotIn(parse("2023-10-01 10:10:10", l1),
		parse("2023-10-01 10:10:11", l1), parse("2023-10-01 10:10:12", l1))))
	assert.True(t, gofn.Head(NotIn(parse("2023-10-01", l2), parse("2023-10-02", l2))))
	assert.True(t, gofn.Head(NotIn(parse("2023-10-01", l2),
		parse("2023-10-02", l2), parse("2023-10-03", l2))))

	assert.False(t, gofn.Head(NotIn(parse("2023-10-01 10:10:10", l1), parse("2023-10-01 10:10:10", l1))))
	assert.False(t, gofn.Head(NotIn(parse("2023-10-01", l2), parse("2023-10-01", l2))))
	assert.False(t, gofn.Head(NotIn(parse("2023-10-01", l2),
		parse("2023-10-02", l2), parse("2023-10-01", l2))))
}
