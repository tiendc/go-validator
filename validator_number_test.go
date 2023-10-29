package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_NumEQ(t *testing.T) {
	errs := NumEQ(gofn.New(10), 10).Exec()
	assert.Equal(t, 0, len(errs))
	errs = NumEQ(gofn.New(10), 9).Exec()
	assert.Equal(t, 1, len(errs))
}

func Test_NumGT_NumGTE(t *testing.T) {
	errs := NumGT(gofn.New(10), 9).Exec()
	assert.Equal(t, 0, len(errs))
	errs = NumGTE(gofn.New(10), 10).Exec()
	assert.Equal(t, 0, len(errs))

	errs = NumGT(gofn.New(10), 10).Exec()
	assert.Equal(t, "gt", errs[0].Type())
	errs = NumGTE(gofn.New(10), 11).Exec()
	assert.Equal(t, "gte", errs[0].Type())
}

func Test_NumLT_NumLTE(t *testing.T) {
	errs := NumLT(gofn.New(10), 11).Exec()
	assert.Equal(t, 0, len(errs))
	errs = NumLTE(gofn.New(10), 10).Exec()
	assert.Equal(t, 0, len(errs))

	errs = NumLT(gofn.New(10), 10).Exec()
	assert.Equal(t, "lt", errs[0].Type())
	errs = NumLTE(gofn.New(10), 9).Exec()
	assert.Equal(t, "lte", errs[0].Type())
}

func Test_NumRange(t *testing.T) {
	errs := NumRange(gofn.New(10), 2, 10).Exec()
	assert.Equal(t, 0, len(errs))

	errs = NumRange(gofn.New(10), 1, 9).Exec()
	assert.Equal(t, "range", errs[0].Type())
}

func Test_NumIn(t *testing.T) {
	errs := NumIn(gofn.New(1), 0, 1, 10).Exec()
	assert.Equal(t, 0, len(errs))

	errs = NumIn(gofn.New(1), 0, 2, 10).Exec()
	assert.Equal(t, "in", errs[0].Type())
}

func Test_NumNotIn(t *testing.T) {
	errs := NumNotIn(gofn.New(1), 0, 2, 10).Exec()
	assert.Equal(t, 0, len(errs))

	errs = NumNotIn(gofn.New(1), 0, 1, 2).Exec()
	assert.Equal(t, "not_in", errs[0].Type())
}

func Test_NumDivisibleBy(t *testing.T) {
	errs := NumDivisibleBy(gofn.New(10), 5).Exec()
	assert.Equal(t, 0, len(errs))

	errs = NumDivisibleBy(gofn.New(8), 3).Exec()
	assert.Equal(t, "divisible_by", errs[0].Type())
}

func Test_NumJsSafeInt(t *testing.T) {
	errs := NumJsSafeInt(gofn.New(9007199254740991)).Exec()
	assert.Equal(t, 0, len(errs))

	errs = NumJsSafeInt(gofn.New(9007199254740992)).Exec()
	assert.Equal(t, "js_safe_int", errs[0].Type())
}
