package mapvalidation

import (
	"github.com/tiendc/gofn"

	"github.com/tiendc/go-validator/base"
)

const (
	kLen       = "Len"
	kItemKey   = "ItemKey"
	kItemValue = "ItemValue"
)

func Len[K comparable, V any](m map[K]V, min, max int) (bool, []base.ErrorParam) {
	l := len(m)
	if min <= l && l <= max {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kLen, Value: l}}
}

func KeyIn[K comparable, V any](m map[K]V, vals ...K) (bool, []base.ErrorParam) {
	keys := gofn.MapKeys(m)
	errIdx := base.IsIn(keys, vals)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemKey, Value: keys[errIdx]}}
}

func KeyNotIn[K comparable, V any](m map[K]V, vals ...K) (bool, []base.ErrorParam) {
	keys := gofn.MapKeys(m)
	errIdx := base.IsNotIn(keys, vals)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemKey, Value: keys[errIdx]}}
}

func KeyRange[K base.Number | base.String, V any](m map[K]V, min, max K) (bool, []base.ErrorParam) {
	for k := range m {
		if k < min || k > max {
			return false, []base.ErrorParam{{Key: kItemKey, Value: k}}
		}
	}
	return true, nil
}

func ValueIn[K comparable, V comparable](m map[K]V, vals ...V) (bool, []base.ErrorParam) {
	keys, values := getKeysAndValues(m)
	errIdx := base.IsIn(values, vals)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemKey, Value: keys[errIdx]}, {Key: kItemValue, Value: values[errIdx]}}
}

func ValueNotIn[K comparable, V comparable](m map[K]V, vals ...V) (bool, []base.ErrorParam) {
	keys, values := getKeysAndValues(m)
	errIdx := base.IsNotIn(values, vals)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemKey, Value: keys[errIdx]}, {Key: kItemValue, Value: values[errIdx]}}
}

func ValueRange[K comparable, V base.Number | base.String](m map[K]V, min, max V) (bool, []base.ErrorParam) {
	for k, v := range m {
		if v < min || v > max {
			return false, []base.ErrorParam{{Key: kItemKey, Value: k}, {Key: kItemValue, Value: v}}
		}
	}
	return true, nil
}

func ValueUnique[K comparable, V comparable](m map[K]V) (bool, []base.ErrorParam) {
	keys, values := getKeysAndValues(m)
	errIdx := base.IsUnique(values)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemKey, Value: keys[errIdx]}, {Key: kItemValue, Value: values[errIdx]}}
}

func getKeysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
