package slicevalidation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_EQ(t *testing.T) {
	assert.True(t, gofn.Head(EQ[int](nil, nil)))
	assert.True(t, gofn.Head(EQ([]int{}, []int{})))
	assert.True(t, gofn.Head(EQ([]int{1, 2}, []int{1, 2})))

	ok, params := EQ([]int{1, 2}, []int{1, 3})
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemValue && params[0].Value == 2 &&
		params[1].Key == kItemIndex && params[1].Value == 1)
}

func Test_Len(t *testing.T) {
	assert.True(t, gofn.Head(Len[int](nil, 0, 10)))
	assert.True(t, gofn.Head(Len([]int{}, 0, 10)))
	assert.True(t, gofn.Head(Len([]int{1, 2}, 0, 2)))

	ok, params := Len([]int{1, 2}, 3, 10)
	assert.False(t, ok)
	assert.True(t, params[0].Key == kLen && params[0].Value == 2)
}

func Test_Unique(t *testing.T) {
	assert.True(t, gofn.Head(Unique[int](nil)))
	assert.True(t, gofn.Head(Unique([]int{})))
	assert.True(t, gofn.Head(Unique([]int{1, 2, 3})))

	ok, params := Unique([]int{0, 1, 2, 0})
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemValue && params[0].Value == 0 &&
		params[1].Key == kItemIndex && params[1].Value == 3)
}

func Test_Sorted(t *testing.T) {
	assert.True(t, gofn.Head(Sorted[int](nil)))
	assert.True(t, gofn.Head(Sorted([]int{})))
	assert.True(t, gofn.Head(Sorted([]int{1, 2, 3})))

	ok, params := Sorted([]int{0, 1, 2, -1})
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemValue && params[0].Value == -1 &&
		params[1].Key == kItemIndex && params[1].Value == 3)
}

func Test_SortedDesc(t *testing.T) {
	assert.True(t, gofn.Head(SortedDesc[int](nil)))
	assert.True(t, gofn.Head(SortedDesc([]int{})))
	assert.True(t, gofn.Head(SortedDesc([]int{3, 2, 0, -1})))

	ok, params := SortedDesc([]int{2, -1, 1, 0})
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemValue && params[0].Value == 1 &&
		params[1].Key == kItemIndex && params[1].Value == 2)
}

func Test_ElemIn(t *testing.T) {
	assert.True(t, gofn.Head(ElemIn[int](nil, 0, 1, 2)))
	assert.True(t, gofn.Head(ElemIn([]int{}, 0, 1, 2)))
	assert.True(t, gofn.Head(ElemIn([]int{1, 2}, 0, 1, 2)))

	ok, params := ElemIn([]int{2, 0, 1, 2}, 1, 2)
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemValue && params[0].Value == 0 &&
		params[1].Key == kItemIndex && params[1].Value == 1)
}

func Test_ElemNotIn(t *testing.T) {
	assert.True(t, gofn.Head(ElemNotIn[int](nil, 0, 1, 2)))
	assert.True(t, gofn.Head(ElemNotIn([]int{}, 0, 1, 2)))
	assert.True(t, gofn.Head(ElemNotIn([]int{1, 2})))
	assert.True(t, gofn.Head(ElemNotIn([]int{1, 2}, 0, 3, 4, 5)))

	ok, params := ElemNotIn([]int{0, 1, 2}, 1, 2, 3)
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemValue && params[0].Value == 1 &&
		params[1].Key == kItemIndex && params[1].Value == 1)
}

func Test_ElemRange(t *testing.T) {
	assert.True(t, gofn.Head(ElemRange[int](nil, 0, 10)))
	assert.True(t, gofn.Head(ElemRange([]int{}, 0, 10)))
	assert.True(t, gofn.Head(ElemRange([]int{0, 1, 2, 10}, 0, 10)))

	ok, params := ElemRange([]int{0, 1, 11, 13}, 0, 10)
	assert.False(t, ok)
	assert.True(t, params[0].Key == kItemValue && params[0].Value == 11 &&
		params[1].Key == kItemIndex && params[1].Value == 2)
}
