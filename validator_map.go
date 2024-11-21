package validation

import (
	"github.com/tiendc/go-validator/base"
	mapFunc "github.com/tiendc/go-validator/base/map"
)

const (
	mapType = "map"
)

// Map allows validating every map entry
func Map[K comparable, V any, M ~map[K]V](m M) MapContentValidator[K, V, M] {
	return NewMapContentValidator(m)
}

// MapLen validates the input map must have length in the specified range
func MapLen[K comparable, V any, M ~map[K]V](m M, min, max int) SingleValidator {
	return call3[M]("len", mapType, "Min", "Max", mapFunc.Len[K, V, M])(m, min, max)
}

// MapKeyIn validates the input map must have keys in the specified values
func MapKeyIn[K comparable, V any, M ~map[K]V](m M, keys ...K) SingleValidator {
	return call2N[M]("key_in", mapType, "TargetValue", mapFunc.KeyIn[K, V, M])(m, keys...)
}

// MapKeyNotIn validates the input map must have keys not in the specified values
func MapKeyNotIn[K comparable, V any, M ~map[K]V](m M, keys ...K) SingleValidator {
	return call2N[M]("key_not_in", mapType, "TargetValue", mapFunc.KeyNotIn[K, V, M])(m, keys...)
}

// MapKeyRange validates the input map must have keys in the specified range.
// Only applies to key type of number or string.
func MapKeyRange[K base.Number | base.String, V any, M ~map[K]V](m M, min, max K) SingleValidator {
	return call3[M]("key_range", mapType, "Min", "Max", mapFunc.KeyRange[K, V, M])(m, min, max)
}

// MapValueIn validates the input map must have values in the specified values
func MapValueIn[K comparable, V comparable, M ~map[K]V](m M, values ...V) SingleValidator {
	return call2N[M]("value_in", mapType, "TargetValue", mapFunc.ValueIn[K, V, M])(m, values...)
}

// MapValueNotIn validates the input map must have values not in the specified values
func MapValueNotIn[K comparable, V comparable, M ~map[K]V](m M, values ...V) SingleValidator {
	return call2N[M]("value_not_in", mapType, "TargetValue", mapFunc.ValueNotIn[K, V, M])(m, values...)
}

// MapValueRange validates the input map must have values in the specified range.
// Only applies to value type of number or string.
func MapValueRange[K comparable, V base.Number | base.String, M ~map[K]V](m M, min, max V) SingleValidator {
	return call3[M]("value_range", mapType, "Min", "Max", mapFunc.ValueRange[K, V, M])(m, min, max)
}

// MapValueUnique validates the input map must have unique values
func MapValueUnique[K comparable, V comparable, M ~map[K]V](m M) SingleValidator {
	return call1[M]("value_unique", mapType, mapFunc.ValueUnique[K, V, M])(m)
}
