package mapvalidation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_Len(t *testing.T) {
	assert.True(t, gofn.Head(Len[int, int](map[int]int(nil), 0, 10)))
	assert.True(t, gofn.Head(Len(map[int]int{}, 0, 10)))
	assert.True(t, gofn.Head(Len(map[int]int{1: 1, 2: 2}, 0, 2)))

	ok, params := Len(map[int]int{1: 1, 2: 2}, 3, 10)
	assert.False(t, ok)
	assert.True(t, params[0].Key == kLen && params[0].Value == 2)
}

func Test_KeyIn(t *testing.T) {
	assert.True(t, gofn.Head(KeyIn[int, int](map[int]int(nil), 0, 1, 2)))
	assert.True(t, gofn.Head(KeyIn[int, int](map[int]int(nil))))
	assert.True(t, gofn.Head(KeyIn(map[int]int{}, 0, 1, 2)))
	assert.True(t, gofn.Head(KeyIn(map[int]int{1: 1, 2: 2}, 0, 1, 2)))

	ok, params := KeyIn(map[int]int{1: 1, 2: 2}, 0, 1)
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemKey && params[0].Value == 2)
}

func Test_KeyNotIn(t *testing.T) {
	assert.True(t, gofn.Head(KeyNotIn[int, int](map[int]int(nil), 0, 1, 2)))
	assert.True(t, gofn.Head(KeyNotIn[int, int](map[int]int(nil))))
	assert.True(t, gofn.Head(KeyNotIn(map[int]int{}, 0, 1, 2)))
	assert.True(t, gofn.Head(KeyNotIn(map[int]int{1: 1, 2: 2}, 3, 4, 5)))

	ok, params := KeyNotIn(map[int]int{1: 1, 2: 2}, 0, 1)
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemKey && params[0].Value == 1)
}

func Test_KeyRange(t *testing.T) {
	assert.True(t, gofn.Head(KeyRange[int, int](map[int]int(nil), 0, 10)))
	assert.True(t, gofn.Head(KeyRange[int, int](map[int]int(nil), 0, 10)))
	assert.True(t, gofn.Head(KeyRange(map[int]int{}, 0, 10)))
	assert.True(t, gofn.Head(KeyRange(map[int]int{1: 1, 2: 2}, 0, 10)))

	ok, params := KeyRange(map[int]int{1: 1, 2: 2}, 2, 10)
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemKey && params[0].Value == 1)
}

func Test_ValueIn(t *testing.T) {
	assert.True(t, gofn.Head(ValueIn[int, int](map[int]int(nil), 0, 1, 2)))
	assert.True(t, gofn.Head(ValueIn[int, int](map[int]int(nil))))
	assert.True(t, gofn.Head(ValueIn(map[int]int{}, 0, 1, 2)))
	assert.True(t, gofn.Head(ValueIn(map[int]int{1: 1, 2: 2}, 0, 1, 2)))

	ok, params := ValueIn(map[int]int{1: 1, 2: 2}, 0, 1)
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemKey && params[0].Value == 2 &&
		params[1].Key == kItemValue && params[1].Value == 2)
}

func Test_ValueNotIn(t *testing.T) {
	assert.True(t, gofn.Head(ValueNotIn[int, int](map[int]int(nil), 0, 1, 2)))
	assert.True(t, gofn.Head(ValueNotIn[int, int](map[int]int(nil))))
	assert.True(t, gofn.Head(ValueNotIn(map[int]int{}, 0, 1, 2)))
	assert.True(t, gofn.Head(ValueNotIn(map[int]int{1: 1, 2: 2}, 3, 4, 5)))

	ok, params := ValueNotIn(map[int]int{1: 1, 2: 2}, 0, 1)
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemKey && params[0].Value == 1 &&
		params[1].Key == kItemValue && params[1].Value == 1)
}

func Test_ValueRange(t *testing.T) {
	assert.True(t, gofn.Head(ValueRange[int, int](map[int]int(nil), 0, 10)))
	assert.True(t, gofn.Head(ValueRange[int, int](map[int]int(nil), 0, 10)))
	assert.True(t, gofn.Head(ValueRange(map[int]int{}, 0, 10)))
	assert.True(t, gofn.Head(ValueRange(map[int]int{1: 1, 2: 2}, 0, 10)))

	ok, params := ValueRange(map[int]int{1: 1, 2: 2}, 2, 10)
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemKey && params[0].Value == 1 &&
		params[1].Key == kItemValue && params[1].Value == 1)
}

func Test_ValueUnique(t *testing.T) {
	assert.True(t, gofn.Head(ValueUnique[int, int](map[int]int(nil))))
	assert.True(t, gofn.Head(ValueUnique[int, int](map[int]int(nil))))
	assert.True(t, gofn.Head(ValueUnique(map[int]int{})))
	assert.True(t, gofn.Head(ValueUnique(map[int]int{1: 1, 2: 2})))

	ok, params := ValueUnique(map[int]int{1: 1, 2: 1})
	assert.False(t, ok)
	assert.True(t, params[1].Key == kItemValue && params[1].Value == 1)
}
