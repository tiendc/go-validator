package validation

import (
	"fmt"
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

func Test_SliceUniqueBy(t *testing.T) {
	errs := SliceUniqueBy([]any{1, 2, 3}, func(v any) int { return v.(int) }).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceUniqueBy([]any{1, 2, 3, 1}, func(v any) int { return v.(int) }).Validate(ctxBg)
	assert.Equal(t, "unique", errs[0].Type())
}

func Test_SliceSorted(t *testing.T) {
	errs := SliceSorted([]int{1, 2, 3}).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceSorted([]int{1, 3, 2}).Validate(ctxBg)
	assert.Equal(t, "sorted", errs[0].Type())
}

func Test_SliceSortedBy(t *testing.T) {
	errs := SliceSortedBy([]any{1, 2, 3}, func(v any) int { return v.(int) }).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceSortedBy([]any{1, 3, 2}, func(v any) int { return v.(int) }).Validate(ctxBg)
	assert.Equal(t, "sorted", errs[0].Type())
}

func Test_SliceSortedDesc(t *testing.T) {
	errs := SliceSortedDesc([]int{3, 2, 1}).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceSortedDesc([]int{3, 2, 3}).Validate(ctxBg)
	assert.Equal(t, "sorted_desc", errs[0].Type())
}

func Test_SliceSortedDescBy(t *testing.T) {
	errs := SliceSortedDescBy([]any{3, 2, 1}, func(v any) int { return v.(int) }).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = SliceSortedDescBy([]any{3, 2, 3}, func(v any) int { return v.(int) }).Validate(ctxBg)
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

func Test_SliceElemValidate(t *testing.T) {
	t.Run("nil/empty slice", func(t *testing.T) {
		// Nil slice
		errs := Slice([]string(nil)).ForEach(func(elem string, index int, validator ItemValidator) {
			validator.Validate(StrLen(&elem, 1, 10))
		}).Validate(ctxBg)
		assert.Equal(t, 0, len(errs))

		// Empty slice
		errs = Slice([]string{}).ForEach(func(elem string, index int, validator ItemValidator) {
			validator.Validate(StrLen(&elem, 1, 10))
		}).Validate(ctxBg)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("validate elements", func(t *testing.T) {
		// Validate element
		errs := Slice([]int{3, 2, 1}).ForEach(func(elem int, index int, validator ItemValidator) {
			validator.Validate(
				NumGTE(&elem, 1),
			)
		}).Validate(ctxBg)
		assert.Equal(t, 0, len(errs))

		// Validate element with errors
		errs = Slice([]int{3, 2, 1}).ForEach(func(elem int, index int, validator ItemValidator) {
			validator.Validate(
				NumGT(&elem, 2).OnError(
					SetField(fmt.Sprintf("slice[%d]", index), nil),
					SetCustomKey("ERR_VLD_SLICE_ELEMENT_INVALID"),
				),
			)
		}).OnError().Validate(ctxBg)
		assert.Equal(t, 2, len(errs))
		assert.Equal(t, "gt", errs[0].Type())
		assert.Equal(t, "gt", errs[1].Type())
	})

	t.Run("validate element with Group", func(t *testing.T) {
		errs := Slice([]int{3, 2, 1}).ForEach(func(elem int, index int, validator ItemValidator) {
			validator.Group(
				NumGTE(&elem, 1),
				NumLTE(&elem, 2),
			)
		}).Validate(ctxBg)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "group", errs[0].Type())
	})

	t.Run("validate element with OneOf", func(t *testing.T) {
		errs := Slice([]int{3, 2, 1}).ForEach(func(elem int, index int, validator ItemValidator) {
			validator.OneOf(
				NumLTE(&elem, 2),
				NumGTE(&elem, 1),
			)
		}).Validate(ctxBg)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("validate element with ExactOneOf", func(t *testing.T) {
		errs := Slice([]int{3, 2, 1}).ForEach(func(elem int, index int, validator ItemValidator) {
			validator.ExactOneOf(
				NumGTE(&elem, 1),
				NumLTE(&elem, 2),
			)
		}).Validate(ctxBg)
		assert.Equal(t, 2, len(errs)) // slice[1] and slice[2] not satisfy
	})

	t.Run("validate element with NotOf", func(t *testing.T) {
		errs := Slice([]int{3, 2, 1}).ForEach(func(elem int, index int, validator ItemValidator) {
			validator.NotOf(
				NumGTE(&elem, 1),
				NumLTE(&elem, 2),
			)
		}).Validate(ctxBg)
		assert.Equal(t, 3, len(errs))
	})
}
