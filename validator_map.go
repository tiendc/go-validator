package validation

import (
	"github.com/tiendc/go-validator/base"
	mapFunc "github.com/tiendc/go-validator/base/map"
)

const (
	mapType = "map"
)

// MapLen validates the input map must have length in the specified range
func MapLen[K comparable, V any](m map[K]V, min, max int) SingleValidator {
	return call3[map[K]V]("len", mapType, "Min", "Max", mapFunc.Len[K, V])(m, min, max)
}

// MapKeyIn validates the input map must have keys in the specified values
func MapKeyIn[K comparable, V any](m map[K]V, keys ...K) SingleValidator {
	return call2N[map[K]V]("key_in", mapType, "TargetValue", mapFunc.KeyIn[K, V])(m, keys...)
}

// MapKeyNotIn validates the input map must have keys not in the specified values
func MapKeyNotIn[K comparable, V any](m map[K]V, keys ...K) SingleValidator {
	return call2N[map[K]V]("key_not_in", mapType, "TargetValue", mapFunc.KeyNotIn[K, V])(m, keys...)
}

// MapKeyRange validates the input map must have keys in the specified range.
// Only applies to key type of number or string.
func MapKeyRange[K base.Number | base.String, V any](m map[K]V, min, max K) SingleValidator {
	return call3[map[K]V]("key_range", mapType, "Min", "Max", mapFunc.KeyRange[K, V])(m, min, max)
}

// MapValueIn validates the input map must have values in the specified values
func MapValueIn[K comparable, V comparable](m map[K]V, values ...V) SingleValidator {
	return call2N[map[K]V]("value_in", mapType, "TargetValue", mapFunc.ValueIn[K, V])(m, values...)
}

// MapValueNotIn validates the input map must have values not in the specified values
func MapValueNotIn[K comparable, V comparable](m map[K]V, values ...V) SingleValidator {
	return call2N[map[K]V]("value_not_in", mapType, "TargetValue", mapFunc.ValueNotIn[K, V])(m, values...)
}

// MapValueRange validates the input map must have values in the specified range.
// Only applies to value type of number or string.
func MapValueRange[K comparable, V base.Number | base.String](m map[K]V, min, max V) SingleValidator {
	return call3[map[K]V]("value_range", mapType, "Min", "Max", mapFunc.ValueRange[K, V])(m, min, max)
}

// MapValueUnique validates the input map must have unique values
func MapValueUnique[K comparable, V comparable](m map[K]V) SingleValidator {
	return call1[map[K]V]("value_unique", mapType, mapFunc.ValueUnique[K, V])(m)
}
