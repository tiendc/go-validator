package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_Nil_NotNil(t *testing.T) {
	// Success cases
	errsNil := Nil[any](nil).Exec()
	errsNotNil := NotNil[any](nil).Exec()
	assert.True(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	var pStr *string
	errsNil = Nil(pStr).Exec()
	errsNotNil = NotNil(pStr).Exec()
	assert.True(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	var pSlice *[]string
	errsNil = Nil(pSlice).Exec()
	errsNotNil = NotNil(pSlice).Exec()
	assert.True(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	var pMap *map[int]bool
	errsNil = Nil(pMap).Exec()
	errsNotNil = NotNil(pMap).Exec()
	assert.True(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	// Failure cases
	errsNil = Nil(gofn.New[any]("")).Exec()
	errsNotNil = NotNil(gofn.New[any]("")).Exec()
	assert.False(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	errsNil = Nil(gofn.New[int](0)).Exec()
	errsNotNil = NotNil(gofn.New[int](0)).Exec()
	assert.False(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	var aSlice []string
	errsNil = Nil(&aSlice).Exec()
	errsNotNil = NotNil(&aSlice).Exec()
	assert.False(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	var aMap map[int]bool
	errsNil = Nil(&aMap).Exec()
	errsNotNil = NotNil(&aMap).Exec()
	assert.False(t, len(errsNil) == 0 && len(errsNotNil) == 1)
}

func Test_Required(t *testing.T) {
	// Failure cases
	errs := Required(nil).Exec()
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	aInt := int32(0)
	errs = Required(&aInt).Exec()
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	aStr := ""
	errs = Required(&aStr).Exec()
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	aSlice := []string{}
	errs = Required(&aSlice).Exec()
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	aMap := map[int]bool{}
	errs = Required(&aMap).Exec()
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	var aAny any
	aAny = aMap
	errs = Required(&aAny).Exec()
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	// Success cases
	aInt = 1
	errs = Required(&aInt).Exec()
	assert.Equal(t, 0, len(errs))

	aStr = "a"
	errs = Required(&aStr).Exec()
	assert.Equal(t, 0, len(errs))

	aSlice = []string{"a", "b"}
	errs = Required(&aSlice).Exec()
	assert.Equal(t, 0, len(errs))

	aMap = map[int]bool{0: false}
	errs = Required(&aMap).Exec()
	assert.Equal(t, 0, len(errs))

	aAny = aSlice
	errs = Required(&aAny).Exec()
	assert.Equal(t, 0, len(errs))
}
