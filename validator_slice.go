package validation

import (
	"github.com/tiendc/go-validator/base"
	sliceFunc "github.com/tiendc/go-validator/base/slice"
)

const (
	sliceType = "slice"
)

// SliceLen validates the input slice must have length in the specified range
func SliceLen[T any](v []T, min, max int) SingleValidator {
	return call3[[]T]("len", sliceType, "Min", "Max", sliceFunc.Len[T])(v, min, max)
}

// SliceUnique validates the input slice must contain only unique items
func SliceUnique[T comparable](v []T) SingleValidator {
	return call1[[]T]("unique", sliceType, sliceFunc.Unique[T])(v)
}

// SliceSorted validates the input slice must be sorted in ascending order
func SliceSorted[T base.Number | base.String](v []T) SingleValidator {
	return call1[[]T]("sorted", sliceType, sliceFunc.Sorted[T])(v)
}

// SliceSortedDesc validates the input slice must be sorted in descending order
func SliceSortedDesc[T base.Number | base.String](v []T) SingleValidator {
	return call1[[]T]("sorted_desc", sliceType, sliceFunc.SortedDesc[T])(v)
}

// SliceElemIn validates the input slice must contain items in the specified values
func SliceElemIn[T comparable](v []T, list ...T) SingleValidator {
	return call2N[[]T]("elem_in", sliceType, "TargetValue", sliceFunc.ElemIn[T])(v, list...)
}

// SliceElemNotIn validates the input slice must contain items not in the specified values
func SliceElemNotIn[T comparable](v []T, list ...T) SingleValidator {
	return call2N[[]T]("elem_not_in", sliceType, "TargetValue", sliceFunc.ElemNotIn[T])(v, list...)
}

// SliceElemRange validates the input slice must contain items in the specified range
func SliceElemRange[T base.Number | base.String](v []T, min, max T) SingleValidator {
	return call3[[]T]("elem_range", sliceType, "Min", "Max", sliceFunc.ElemRange[T])(v, min, max)
}
