package slicevalidation

import (
	"github.com/tiendc/go-validator/base"
)

const (
	kLen       = "Len"
	kItemValue = "ItemValue"
	kItemIndex = "ItemIndex"
)

// EQ checks a slice must equal to a slice
func EQ[T comparable, S ~[]T](s S, s2 S) (bool, []base.ErrorParam) {
	if len(s) != len(s2) {
		return false, nil
	}
	for i, v := range s {
		if s[i] != s2[i] {
			return false, []base.ErrorParam{{Key: kItemValue, Value: v}, {Key: kItemIndex, Value: i}}
		}
	}
	return true, nil
}

// Len checks slice length must be in a range
func Len[T any, S ~[]T](s S, min, max int) (bool, []base.ErrorParam) {
	l := len(s)
	if min <= l && l <= max {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kLen, Value: l}}
}

// Unique checks slice items must be unique
func Unique[T comparable, S ~[]T](s S) (bool, []base.ErrorParam) {
	errIdx := base.IsUnique(s)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemValue, Value: s[errIdx]}, {Key: kItemIndex, Value: errIdx}}
}

// UniqueBy checks slice items must be unique
func UniqueBy[T any, U comparable, S ~[]T](s S, keyFn func(T) U) (bool, []base.ErrorParam) {
	errIdx := base.IsUniqueBy(s, keyFn)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemValue, Value: s[errIdx]}, {Key: kItemIndex, Value: errIdx}}
}

// Sorted checks slice items must be in ascending order
func Sorted[T base.Number | base.String, S ~[]T](s S) (bool, []base.ErrorParam) {
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return false, []base.ErrorParam{{Key: kItemValue, Value: s[i]}, {Key: kItemIndex, Value: i}}
		}
	}
	return true, nil
}

// SortedBy checks slice items must be in ascending order defined by the function
func SortedBy[T any, U base.Number | base.String, S ~[]T](s S, keyFn func(T) U) (bool, []base.ErrorParam) {
	length := len(s)
	if length <= 1 {
		return true, nil
	}
	prevVal := keyFn(s[0])
	for i := 1; i < length; i++ {
		currVal := keyFn(s[i])
		if prevVal > currVal {
			return false, []base.ErrorParam{{Key: kItemValue, Value: s[i]}, {Key: kItemIndex, Value: i}}
		}
		prevVal = currVal
	}
	return true, nil
}

// SortedDesc checks slice items must be in descending order
func SortedDesc[T base.Number | base.String, S ~[]T](s S) (bool, []base.ErrorParam) {
	for i := 1; i < len(s); i++ {
		if s[i-1] < s[i] {
			return false, []base.ErrorParam{{Key: kItemValue, Value: s[i]}, {Key: kItemIndex, Value: i}}
		}
	}
	return true, nil
}

// SortedDescBy checks slice items must be in descending order defined by the function
func SortedDescBy[T any, U base.Number | base.String, S ~[]T](s S, keyFn func(T) U) (bool, []base.ErrorParam) {
	length := len(s)
	if length <= 1 {
		return true, nil
	}
	prevVal := keyFn(s[0])
	for i := 1; i < length; i++ {
		currVal := keyFn(s[i])
		if prevVal < currVal {
			return false, []base.ErrorParam{{Key: kItemValue, Value: s[i]}, {Key: kItemIndex, Value: i}}
		}
		prevVal = currVal
	}
	return true, nil
}

// HasElem checks slice must contain the specified values
func HasElem[T comparable, S ~[]T](s S, values ...T) (bool, []base.ErrorParam) {
	elemMap := base.ToMap(s)
	for _, v := range values {
		if _, exists := elemMap[v]; !exists {
			return false, []base.ErrorParam{{Key: kItemValue, Value: v}}
		}
	}
	return true, nil
}

// HasElemBy checks slice must contain values using custom function
func HasElemBy[T any, S ~[]T](s S, isExistFn func(T) bool) (bool, []base.ErrorParam) {
	for i := range s {
		if !isExistFn(s[i]) {
			return false, []base.ErrorParam{{Key: kItemIndex, Value: i}}
		}
	}
	return true, nil
}

// NotHaveElem checks slice must not contain the specified values
func NotHaveElem[T comparable, S ~[]T](s S, values ...T) (bool, []base.ErrorParam) {
	elemMap := base.ToMap(s)
	for _, v := range values {
		if _, exists := elemMap[v]; exists {
			return false, []base.ErrorParam{{Key: kItemValue, Value: v}}
		}
	}
	return true, nil
}

// NotHaveElemBy checks slice must not contain values using custom function
func NotHaveElemBy[T any, S ~[]T](s S, isExistFn func(T) bool) (bool, []base.ErrorParam) {
	for i := range s {
		if isExistFn(s[i]) {
			return false, []base.ErrorParam{{Key: kItemIndex, Value: i}}
		}
	}
	return true, nil
}

// ElemIn checks slice items must be in a list
func ElemIn[T comparable, S ~[]T](s S, values ...T) (bool, []base.ErrorParam) {
	errIdx := base.IsIn(s, values)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemValue, Value: s[errIdx]}, {Key: kItemIndex, Value: errIdx}}
}

// ElemNotIn checks slice items must be not in a list
func ElemNotIn[T comparable, S ~[]T](s S, values ...T) (bool, []base.ErrorParam) {
	errIdx := base.IsNotIn(s, values)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemValue, Value: s[errIdx]}, {Key: kItemIndex, Value: errIdx}}
}

// ElemRange checks slice items must be in a range (applies to item type string or number only)
func ElemRange[T base.Number | base.String, S ~[]T](s S, min, max T) (bool, []base.ErrorParam) {
	for i, v := range s {
		if v < min || v > max {
			return false, []base.ErrorParam{{Key: kItemValue, Value: v}, {Key: kItemIndex, Value: i}}
		}
	}
	return true, nil
}
