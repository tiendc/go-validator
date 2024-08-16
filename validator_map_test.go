package validation

import (
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
