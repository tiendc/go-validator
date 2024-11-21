package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapLen(t *testing.T) {
	errs := MapLen(map[int]int{1: 1, 2: 2}, 2, 10).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = MapLen(map[int]int{1: 1, 2: 2}, 3, 10).Validate(ctxBg)
	assert.Equal(t, "len", errs[0].Type())
}

func Test_MapKeyIn(t *testing.T) {
	errs := MapKeyIn(map[int]int{3: 3, 2: 2, 1: 1}, 1, 2, 3, 4, 5).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = MapKeyIn(map[int]int{3: 3, 2: 2, 1: 1}, 1, 2).Validate(ctxBg)
	assert.Equal(t, "key_in", errs[0].Type())
}

func Test_MapKeyNotIn(t *testing.T) {
	errs := MapKeyNotIn(map[int]int{3: 3, 2: 2, 1: 1}, 4, 5, 6).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = MapKeyNotIn(map[int]int{3: 3, 2: 2, 1: 1}, 1, 2).Validate(ctxBg)
	assert.Equal(t, "key_not_in", errs[0].Type())
}

func Test_MapKeyRange(t *testing.T) {
	errs := MapKeyRange(map[int]int{3: 3, 2: 2, 1: 1}, 0, 3).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = MapKeyRange(map[int]int{3: 3, 2: 2, 1: 1}, 1, 2).Validate(ctxBg)
	assert.Equal(t, "key_range", errs[0].Type())
}

func Test_MapValueIn(t *testing.T) {
	errs := MapValueIn(map[int]int{3: 3, 2: 2, 1: 1}, 1, 2, 3, 4, 5).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = MapValueIn(map[int]int{3: 3, 2: 2, 1: 1}, 1, 2).Validate(ctxBg)
	assert.Equal(t, "value_in", errs[0].Type())
}

func Test_MapValueNotIn(t *testing.T) {
	errs := MapValueNotIn(map[int]int{3: 3, 2: 2, 1: 1}, 4, 5, 6).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = MapValueNotIn(map[int]int{3: 3, 2: 2, 1: 1}, 1, 2).Validate(ctxBg)
	assert.Equal(t, "value_not_in", errs[0].Type())
}

func Test_MapValueRange(t *testing.T) {
	errs := MapValueRange(map[int]int{3: 3, 2: 2, 1: 1}, 0, 3).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = MapValueRange(map[int]int{3: 3, 2: 2, 1: 1}, 1, 2).Validate(ctxBg)
	assert.Equal(t, "value_range", errs[0].Type())
}

func Test_MapValueUnique(t *testing.T) {
	errs := MapValueUnique(map[int]int{3: 3, 2: 2, 1: 1}).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = MapValueUnique(map[int]int{3: 1, 2: 2, 1: 1}).Validate(ctxBg)
	assert.Equal(t, "value_unique", errs[0].Type())
}

func Test_MapContent_Validate(t *testing.T) {
	t.Run("nil/empty map", func(t *testing.T) {
		// Nil map
		errs := Map(map[int]string(nil)).ForEach(func(k int, v string, validator ItemValidator) {
			validator.Validate(StrLen(&v, 1, 10))
		}).Validate(ctxBg)
		assert.Equal(t, 0, len(errs))

		// Empty map
		errs = Map(map[int]string{}).ForEach(func(k int, v string, validator ItemValidator) {
			validator.Validate(StrLen(&v, 1, 10))
		}).Validate(ctxBg)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("validate entries", func(t *testing.T) {
		// Validate entries
		errs := Map(map[int]int{3: 3, 2: 2, 1: 1}).ForEach(func(k int, v int, validator ItemValidator) {
			validator.Validate(
				NumGTE(&v, 1),
			)
		}).Validate(ctxBg)
		assert.Equal(t, 0, len(errs))

		// Validate entries with errors
		errs = Map(map[int]int{3: 3, 2: 2, 1: 1}).ForEach(func(k int, v int, validator ItemValidator) {
			validator.Validate(
				NumGT(&v, 2).OnError(
					SetField(fmt.Sprintf("map[%d]", k), nil),
					SetCustomKey("ERR_VLD_MAP_ENTRY_INVALID"),
				),
			)
		}).OnError().Validate(ctxBg)
		assert.Equal(t, 2, len(errs))
		assert.Equal(t, "gt", errs[0].Type())
		assert.Equal(t, "gt", errs[1].Type())
	})

	t.Run("validate entries with Group", func(t *testing.T) {
		errs := Map(map[int]int{3: 3, 2: 2, 1: 1}).ForEach(func(k int, v int, validator ItemValidator) {
			validator.Group(
				NumGTE(&v, 1),
				NumLTE(&v, 2),
			)
		}).Validate(ctxBg)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "group", errs[0].Type())
	})

	t.Run("validate entries with OneOf", func(t *testing.T) {
		errs := Map(map[int]int{3: 3, 2: 2, 1: 1}).ForEach(func(k int, v int, validator ItemValidator) {
			validator.OneOf(
				NumLTE(&v, 2),
				NumGTE(&v, 1),
			)
		}).Validate(ctxBg)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("validate entries with ExactOneOf", func(t *testing.T) {
		errs := Map(map[int]int{3: 3, 2: 2, 1: 1}).ForEach(func(k int, v int, validator ItemValidator) {
			validator.ExactOneOf(
				NumGTE(&v, 1),
				NumLTE(&v, 2),
			)
		}).Validate(ctxBg)
		assert.Equal(t, 2, len(errs)) // slice[1] and slice[2] not satisfy
	})

	t.Run("validate entries with NotOf", func(t *testing.T) {
		errs := Map(map[int]int{3: 3, 2: 2, 1: 1}).ForEach(func(k int, v int, validator ItemValidator) {
			validator.NotOf(
				NumGTE(&v, 1),
				NumLTE(&v, 2),
			)
		}).Validate(ctxBg)
		assert.Equal(t, 3, len(errs))
	})
}
