package slicevalidation

import (
	"github.com/tiendc/go-validator/base"
)

const (
	kLen       = "Len"
	kItemValue = "ItemValue"
	kItemIndex = "ItemIndex"
)

func EQ[T comparable](s []T, s2 []T) (bool, []base.ErrorParam) {
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

func Len[T any](s []T, min, max int) (bool, []base.ErrorParam) {
	l := len(s)
	if min <= l && l <= max {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kLen, Value: l}}
}

func Unique[T comparable](s []T) (bool, []base.ErrorParam) {
	errIdx := base.IsUnique(s)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemValue, Value: s[errIdx]}, {Key: kItemIndex, Value: errIdx}}
}

func Sorted[T base.Number | base.String](s []T) (bool, []base.ErrorParam) {
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return false, []base.ErrorParam{{Key: kItemValue, Value: s[i]}, {Key: kItemIndex, Value: i}}
		}
	}
	return true, nil
}

func SortedDesc[T base.Number | base.String](s []T) (bool, []base.ErrorParam) {
	for i := 1; i < len(s); i++ {
		if s[i-1] < s[i] {
			return false, []base.ErrorParam{{Key: kItemValue, Value: s[i]}, {Key: kItemIndex, Value: i}}
		}
	}
	return true, nil
}

func ElemIn[T comparable](s []T, values ...T) (bool, []base.ErrorParam) {
	errIdx := base.IsIn(s, values)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemValue, Value: s[errIdx]}, {Key: kItemIndex, Value: errIdx}}
}

func ElemNotIn[T comparable](s []T, values ...T) (bool, []base.ErrorParam) {
	errIdx := base.IsNotIn(s, values)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemValue, Value: s[errIdx]}, {Key: kItemIndex, Value: errIdx}}
}

func ElemRange[T base.Number | base.String](s []T, min, max T) (bool, []base.ErrorParam) {
	for i, v := range s {
		if v < min || v > max {
			return false, []base.ErrorParam{{Key: kItemValue, Value: v}, {Key: kItemIndex, Value: i}}
		}
	}
	return true, nil
}
