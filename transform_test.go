package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_ToLower(t *testing.T) {
	v := ToLower(gofn.New("aBc"))
	assert.Equal(t, "abc", *v)
}

func Test_ToUpper(t *testing.T) {
	v := ToUpper(gofn.New("aBc"))
	assert.Equal(t, "ABC", *v)
}

func Test_ToInt64(t *testing.T) {
	v := ToInt64(gofn.New(123))
	assert.Equal(t, int64(123), *v)
}

func Test_ToUint64(t *testing.T) {
	v := ToUint64(gofn.New(123))
	assert.Equal(t, uint64(123), *v)
}

func Test_ToFloat64(t *testing.T) {
	v := ToFloat64(gofn.New(123.123))
	assert.Equal(t, float64(123.123), *v)
}
