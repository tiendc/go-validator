package validation

import (
	"github.com/tiendc/go-validator/base"
	mapFunc "github.com/tiendc/go-validator/base/map"
)

const (
	mapType = "map"
)

func MapLen[K comparable, V any](m map[K]V, min, max int) SingleValidator {
	return call3[map[K]V]("len", mapType, "Min", "Max", mapFunc.Len[K, V])(m, min, max)
}

func MapKeyIn[K comparable, V any](m map[K]V, keys ...K) SingleValidator {
	return call2N[map[K]V]("key_in", mapType, "TargetValue", mapFunc.KeyIn[K, V])(m, keys...)
}

func MapKeyNotIn[K comparable, V any](m map[K]V, keys ...K) SingleValidator {
	return call2N[map[K]V]("key_not_in", mapType, "TargetValue", mapFunc.KeyNotIn[K, V])(m, keys...)
}

func MapKeyRange[K base.Number | base.String, V any](m map[K]V, min, max K) SingleValidator {
	return call3[map[K]V]("key_range", mapType, "Min", "Max", mapFunc.KeyRange[K, V])(m, min, max)
}

func MapValueIn[K comparable, V comparable](m map[K]V, values ...V) SingleValidator {
	return call2N[map[K]V]("value_in", mapType, "TargetValue", mapFunc.ValueIn[K, V])(m, values...)
}

func MapValueNotIn[K comparable, V comparable](m map[K]V, values ...V) SingleValidator {
	return call2N[map[K]V]("value_not_in", mapType, "TargetValue", mapFunc.ValueNotIn[K, V])(m, values...)
}

func MapValueRange[K comparable, V base.Number | base.String](m map[K]V, min, max V) SingleValidator {
	return call3[map[K]V]("value_range", mapType, "Min", "Max", mapFunc.ValueRange[K, V])(m, min, max)
}

func MapValueUnique[K comparable, V comparable](m map[K]V) SingleValidator {
	return call1[map[K]V]("value_unique", mapType, mapFunc.ValueUnique[K, V])(m)
}
