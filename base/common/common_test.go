package commonvalidation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_Nil_NotNil(t *testing.T) {
	// Success cases
	isNil, _ := Nil[any](nil)
	isNotNil, _ := NotNil[any](nil)
	assert.True(t, isNil && !isNotNil)

	var pStr *string
	isNil, _ = Nil(pStr)
	isNotNil, _ = NotNil(pStr)
	assert.True(t, isNil && !isNotNil)

	var pSlice *[]string
	isNil, _ = Nil(pSlice)
	isNotNil, _ = NotNil(pSlice)
	assert.True(t, isNil && !isNotNil)

	var pMap *map[int]bool
	isNil, _ = Nil(pMap)
	isNotNil, _ = NotNil(pMap)
	assert.True(t, isNil && !isNotNil)

	// Failure cases
	isNil, _ = Nil(gofn.ToPtr[any](""))
	isNotNil, _ = NotNil(gofn.ToPtr[any](""))
	assert.False(t, isNil && !isNotNil)

	isNil, _ = Nil(gofn.ToPtr[int](0))
	isNotNil, _ = NotNil(gofn.ToPtr[int](0))
	assert.False(t, isNil && !isNotNil)

	var aSlice []string
	isNil, _ = Nil(&aSlice)
	isNotNil, _ = NotNil(&aSlice)
	assert.False(t, isNil && !isNotNil)

	var aMap map[int]bool
	isNil, _ = Nil(&aMap)
	isNotNil, _ = NotNil(&aMap)
	assert.False(t, isNil && !isNotNil)
}

func Test_Required(t *testing.T) {
	// Failure cases
	assert.False(t, gofn.Head(Required(nil)))

	aInt := int32(0)
	assert.False(t, gofn.Head(Required(&aInt)))

	aStr := ""
	assert.False(t, gofn.Head(Required(&aStr)))

	aSlice := []string{}
	assert.False(t, gofn.Head(Required(&aSlice)))

	aMap := map[int]bool{}
	assert.False(t, gofn.Head(Required(&aMap)))

	var aAny any
	aAny = aMap
	assert.False(t, gofn.Head(Required(&aAny)))

	// Success cases
	aInt = 1
	assert.True(t, gofn.Head(Required(&aInt)))

	aStr = "a"
	assert.True(t, gofn.Head(Required(&aStr)))

	aSlice = []string{"a", "b"}
	assert.True(t, gofn.Head(Required(&aSlice)))

	aMap = map[int]bool{0: false}
	assert.True(t, gofn.Head(Required(&aMap)))

	aAny = aSlice
	assert.True(t, gofn.Head(Required(&aAny)))
}
