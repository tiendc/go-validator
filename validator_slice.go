package validation

import (
	"github.com/tiendc/go-validator/base"
	sliceFunc "github.com/tiendc/go-validator/base/slice"
)

const (
	sliceType = "slice"
)

// Slice allows validating slice elements
func Slice[T any, S ~[]T](s S) SliceContentValidator[T, S] {
	return NewSliceContentValidator(s)
}

// SliceLen validates the input slice must have length in the specified range
func SliceLen[T any, S ~[]T](v S, min, max int) SingleValidator {
	return call3[S]("len", sliceType, "Min", "Max", sliceFunc.Len[T, S])(v, min, max)
}

// SliceUnique validates the input slice must contain only unique items
func SliceUnique[T comparable, S ~[]T](v S) SingleValidator {
	return call1[S]("unique", sliceType, sliceFunc.Unique[T, S])(v)
}

// SliceSorted validates the input slice must be sorted in ascending order
func SliceSorted[T base.Number | base.String, S ~[]T](v S) SingleValidator {
	return call1[S]("sorted", sliceType, sliceFunc.Sorted[T, S])(v)
}

// SliceSortedDesc validates the input slice must be sorted in descending order
func SliceSortedDesc[T base.Number | base.String, S ~[]T](v S) SingleValidator {
	return call1[S]("sorted_desc", sliceType, sliceFunc.SortedDesc[T, S])(v)
}

// SliceElemIn validates the input slice must contain items in the specified values
func SliceElemIn[T comparable, S ~[]T](v S, list ...T) SingleValidator {
	return call2N[S]("elem_in", sliceType, "TargetValue", sliceFunc.ElemIn[T, S])(v, list...)
}

// SliceElemNotIn validates the input slice must contain items not in the specified values
func SliceElemNotIn[T comparable, S ~[]T](v S, list ...T) SingleValidator {
	return call2N[S]("elem_not_in", sliceType, "TargetValue", sliceFunc.ElemNotIn[T, S])(v, list...)
}

// SliceElemRange validates the input slice must contain items in the specified range
func SliceElemRange[T base.Number | base.String, S ~[]T](v S, min, max T) SingleValidator {
	return call3[S]("elem_range", sliceType, "Min", "Max", sliceFunc.ElemRange[T, S])(v, min, max)
}
