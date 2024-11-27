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

// SliceUniqueBy validates the input slice must contain only unique items
func SliceUniqueBy[T any, U comparable, S ~[]T](v S, keyFn func(T) U) SingleValidator {
	return call2[S]("unique", sliceType, "KeyFunction", sliceFunc.UniqueBy[T, U, S])(v, keyFn)
}

// SliceSorted validates the input slice must be sorted in ascending order
func SliceSorted[T base.Number | base.String, S ~[]T](v S) SingleValidator {
	return call1[S]("sorted", sliceType, sliceFunc.Sorted[T, S])(v)
}

// SliceSortedBy validates the input slice must be sorted in ascending order defined by the key function
func SliceSortedBy[T any, U base.Number | base.String, S ~[]T](v S, keyFn func(T) U) SingleValidator {
	return call2[S]("sorted", sliceType, "KeyFunction", sliceFunc.SortedBy[T, U, S])(v, keyFn)
}

// SliceSortedDesc validates the input slice must be sorted in descending order
func SliceSortedDesc[T base.Number | base.String, S ~[]T](v S) SingleValidator {
	return call1[S]("sorted_desc", sliceType, sliceFunc.SortedDesc[T, S])(v)
}

// SliceSortedDescBy validates the input slice must be sorted in ascending order defined by the key function
func SliceSortedDescBy[T any, U base.Number | base.String, S ~[]T](v S, keyFn func(T) U) SingleValidator {
	return call2[S]("sorted_desc", sliceType, "KeyFunction", sliceFunc.SortedDescBy[T, U, S])(v, keyFn)
}

// SliceHasElem validates the input slice must contain the specified values
func SliceHasElem[T comparable, S ~[]T](v S, list ...T) SingleValidator {
	return call2N[S]("has_elem", sliceType, "TargetValue", sliceFunc.HasElem[T, S])(v, list...)
}

// SliceNotHaveElem validates the input slice must not contain the specified values
func SliceNotHaveElem[T comparable, S ~[]T](v S, list ...T) SingleValidator {
	return call2N[S]("not_have_elem", sliceType, "TargetValue", sliceFunc.NotHaveElem[T, S])(v, list...)
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
