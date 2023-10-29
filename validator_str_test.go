package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_StrLen(t *testing.T) {
	errs := StrLen(gofn.New("chào"), 2, 10).Exec()
	assert.Equal(t, 0, len(errs))

	errs = StrLen(gofn.New("chào"), 1, 3).Exec()
	assert.Equal(t, "len", errs[0].Type())
}

func Test_StrByteLen(t *testing.T) {
	errs := StrByteLen(gofn.New("ab "), 2, 10).Exec()
	assert.Equal(t, 0, len(errs))

	errs = StrByteLen(gofn.New("chào"), 1, 4).Exec()
	assert.Equal(t, "byte_len", errs[0].Type())
}

func Test_StrEQ(t *testing.T) {
	errs := NumEQ(gofn.New(10), 10).Exec()
	assert.Equal(t, 0, len(errs))
	errs = NumEQ(gofn.New(10), 9).Exec()
	assert.Equal(t, 1, len(errs))
}

func Test_StrIn(t *testing.T) {
	errs := StrIn(gofn.New("chào"), "", "a", "chào").Exec()
	assert.Equal(t, 0, len(errs))

	errs = StrIn(gofn.New("a"), "", "b").Exec()
	assert.Equal(t, "in", errs[0].Type())
}

func Test_StrNotIn(t *testing.T) {
	errs := StrNotIn(gofn.New("chào"), "", "a", "b").Exec()
	assert.Equal(t, 0, len(errs))

	errs = StrNotIn(gofn.New("a"), "", "a", "b").Exec()
	assert.Equal(t, "not_in", errs[0].Type())
}
