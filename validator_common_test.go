package validation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

var ctxBg = context.Background()

func Test_Nil_NotNil(t *testing.T) {
	// Success cases
	errsNil := Nil[any](nil).Validate(ctxBg)
	errsNotNil := NotNil[any](nil).Validate(ctxBg)
	assert.True(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	var pStr *string
	errsNil = Nil(pStr).Validate(ctxBg)
	errsNotNil = NotNil(pStr).Validate(ctxBg)
	assert.True(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	var pSlice *[]string
	errsNil = Nil(pSlice).Validate(ctxBg)
	errsNotNil = NotNil(pSlice).Validate(ctxBg)
	assert.True(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	var pMap *map[int]bool
	errsNil = Nil(pMap).Validate(ctxBg)
	errsNotNil = NotNil(pMap).Validate(ctxBg)
	assert.True(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	// Failure cases
	errsNil = Nil(gofn.ToPtr[any]("")).Validate(ctxBg)
	errsNotNil = NotNil(gofn.ToPtr[any]("")).Validate(ctxBg)
	assert.False(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	errsNil = Nil(gofn.ToPtr[int](0)).Validate(ctxBg)
	errsNotNil = NotNil(gofn.ToPtr[int](0)).Validate(ctxBg)
	assert.False(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	var aSlice []string
	errsNil = Nil(&aSlice).Validate(ctxBg)
	errsNotNil = NotNil(&aSlice).Validate(ctxBg)
	assert.False(t, len(errsNil) == 0 && len(errsNotNil) == 1)

	var aMap map[int]bool
	errsNil = Nil(&aMap).Validate(ctxBg)
	errsNotNil = NotNil(&aMap).Validate(ctxBg)
	assert.False(t, len(errsNil) == 0 && len(errsNotNil) == 1)
}

func Test_Required(t *testing.T) {
	// Failure cases
	errs := Required(nil).Validate(ctxBg)
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	aInt := int32(0)
	errs = Required(&aInt).Validate(ctxBg)
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	aStr := ""
	errs = Required(&aStr).Validate(ctxBg)
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	aSlice := []string{}
	errs = Required(&aSlice).Validate(ctxBg)
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	aMap := map[int]bool{}
	errs = Required(&aMap).Validate(ctxBg)
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	var aAny any
	aAny = aMap
	errs = Required(&aAny).Validate(ctxBg)
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "required", errs[0].Type())

	// Success cases
	aInt = 1
	errs = Required(&aInt).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	aStr = "a"
	errs = Required(&aStr).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	aSlice = []string{"a", "b"}
	errs = Required(&aSlice).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	aMap = map[int]bool{0: false}
	errs = Required(&aMap).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	aAny = aSlice
	errs = Required(&aAny).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))
}
