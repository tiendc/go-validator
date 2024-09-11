package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_NumEQ(t *testing.T) {
	errs := NumEQ(gofn.ToPtr(10), 10).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))
	errs = NumEQ(gofn.ToPtr(10), 9).Validate(ctxBg)
	assert.Equal(t, 1, len(errs))
}

func Test_NumGT_NumGTE(t *testing.T) {
	errs := NumGT(gofn.ToPtr(10), 9).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))
	errs = NumGTE(gofn.ToPtr(10), 10).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = NumGT(gofn.ToPtr(10), 10).Validate(ctxBg)
	assert.Equal(t, "gt", errs[0].Type())
	errs = NumGTE(gofn.ToPtr(10), 11).Validate(ctxBg)
	assert.Equal(t, "gte", errs[0].Type())
}

func Test_NumLT_NumLTE(t *testing.T) {
	errs := NumLT(gofn.ToPtr(10), 11).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))
	errs = NumLTE(gofn.ToPtr(10), 10).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = NumLT(gofn.ToPtr(10), 10).Validate(ctxBg)
	assert.Equal(t, "lt", errs[0].Type())
	errs = NumLTE(gofn.ToPtr(10), 9).Validate(ctxBg)
	assert.Equal(t, "lte", errs[0].Type())
}

func Test_NumRange(t *testing.T) {
	errs := NumRange(gofn.ToPtr(10), 2, 10).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = NumRange(gofn.ToPtr(10), 1, 9).Validate(ctxBg)
	assert.Equal(t, "range", errs[0].Type())
}

func Test_NumIn(t *testing.T) {
	errs := NumIn(gofn.ToPtr(1), 0, 1, 10).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = NumIn(gofn.ToPtr(1), 0, 2, 10).Validate(ctxBg)
	assert.Equal(t, "in", errs[0].Type())
}

func Test_NumNotIn(t *testing.T) {
	errs := NumNotIn(gofn.ToPtr(1), 0, 2, 10).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = NumNotIn(gofn.ToPtr(1), 0, 1, 2).Validate(ctxBg)
	assert.Equal(t, "not_in", errs[0].Type())
}

func Test_NumDivisibleBy(t *testing.T) {
	errs := NumDivisibleBy(gofn.ToPtr(10), 5).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = NumDivisibleBy(gofn.ToPtr(8), 3).Validate(ctxBg)
	assert.Equal(t, "divisible_by", errs[0].Type())
}

func Test_NumJsSafeInt(t *testing.T) {
	errs := NumJsSafeInt(gofn.ToPtr(9007199254740991)).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = NumJsSafeInt(gofn.ToPtr(9007199254740992)).Validate(ctxBg)
	assert.Equal(t, "js_safe_int", errs[0].Type())
}
