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

// Len checks map size must be in a range
func Len[K comparable, V any, M ~map[K]V](m M, min, max int) (bool, []base.ErrorParam) {
	l := len(m)
	if min <= l && l <= max {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kLen, Value: l}}
}

// KeyIn checks map keys must be in a list
func KeyIn[K comparable, V any, M ~map[K]V](m M, vals ...K) (bool, []base.ErrorParam) {
	keys := gofn.MapKeys(m)
	errIdx := base.IsIn(keys, vals)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemKey, Value: keys[errIdx]}}
}

// KeyNotIn checks map keys must be not in a list
func KeyNotIn[K comparable, V any, M ~map[K]V](m M, vals ...K) (bool, []base.ErrorParam) {
	keys := gofn.MapKeys(m)
	errIdx := base.IsNotIn(keys, vals)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemKey, Value: keys[errIdx]}}
}

// KeyRange checks map keys must be in a range (applies to key type string and number only)
func KeyRange[K base.Number | base.String, V any, M ~map[K]V](m M, min, max K) (bool, []base.ErrorParam) {
	for k := range m {
		if k < min || k > max {
			return false, []base.ErrorParam{{Key: kItemKey, Value: k}}
		}
	}
	return true, nil
}

// ValueIn checks map values must be in a list
func ValueIn[K comparable, V comparable, M ~map[K]V](m M, vals ...V) (bool, []base.ErrorParam) {
	keys, values := getKeysAndValues(m)
	errIdx := base.IsIn(values, vals)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemKey, Value: keys[errIdx]}, {Key: kItemValue, Value: values[errIdx]}}
}

// ValueNotIn checks map values must be not in a list
func ValueNotIn[K comparable, V comparable, M ~map[K]V](m M, vals ...V) (bool, []base.ErrorParam) {
	keys, values := getKeysAndValues(m)
	errIdx := base.IsNotIn(values, vals)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemKey, Value: keys[errIdx]}, {Key: kItemValue, Value: values[errIdx]}}
}

// ValueRange checks map values must be in a range (applies to value type string and number only)
func ValueRange[K comparable, V base.Number | base.String, M ~map[K]V](m M, min, max V) (bool, []base.ErrorParam) {
	for k, v := range m {
		if v < min || v > max {
			return false, []base.ErrorParam{{Key: kItemKey, Value: k}, {Key: kItemValue, Value: v}}
		}
	}
	return true, nil
}

// ValueUnique checks map values must be unique
func ValueUnique[K comparable, V comparable, M ~map[K]V](m M) (bool, []base.ErrorParam) {
	keys, values := getKeysAndValues(m)
	errIdx := base.IsUnique(values)
	if errIdx == -1 {
		return true, nil
	}
	return false, []base.ErrorParam{{Key: kItemKey, Value: keys[errIdx]}, {Key: kItemValue, Value: values[errIdx]}}
}

func getKeysAndValues[K comparable, V any, M ~map[K]V](m M) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
