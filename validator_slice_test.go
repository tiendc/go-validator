package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SliceLen(t *testing.T) {
	errs := SliceLen([]int{1, 2}, 2, 10).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceLen([]int{1, 2}, 3, 10).Validate(ctxBg)
	assert.Equal(t, "len", errs[0].Type())
}

func Test_SliceUnique(t *testing.T) {
	errs := SliceUnique([]int{1, 2, 3}).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceUnique([]int{1, 2, 3, 1}).Validate(ctxBg)
	assert.Equal(t, "unique", errs[0].Type())
}

func Test_SliceSorted(t *testing.T) {
	errs := SliceSorted([]int{1, 2, 3}).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceSorted([]int{1, 3, 2}).Validate(ctxBg)
	assert.Equal(t, "sorted", errs[0].Type())
}

func Test_SliceSortedDesc(t *testing.T) {
	errs := SliceSortedDesc([]int{3, 2, 1}).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceSortedDesc([]int{3, 2, 3}).Validate(ctxBg)
	assert.Equal(t, "sorted_desc", errs[0].Type())
}

func Test_SliceElemIn(t *testing.T) {
	errs := SliceElemIn([]int{3, 2, 1}, 1, 2, 3, 4, 5).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceElemIn([]int{3, 2, 1}, 1, 2).Validate(ctxBg)
	assert.Equal(t, "elem_in", errs[0].Type())
}

func Test_SliceElemNotIn(t *testing.T) {
	errs := SliceElemNotIn([]int{3, 2, 1}, 4, 5, 6).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceElemNotIn([]int{3, 2, 1}, 1, 4, 5).Validate(ctxBg)
	assert.Equal(t, "elem_not_in", errs[0].Type())
}

func Test_SliceElemRange(t *testing.T) {
	errs := SliceElemRange([]int{3, 2, 1}, 0, 3).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceElemRange([]int{3, 2, 1}, 1, 2).Validate(ctxBg)
	assert.Equal(t, "elem_range", errs[0].Type())
}
