package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_ToLower(t *testing.T) {
	assert.Nil(t, ToLower[string](nil))
	v := ToLower(gofn.ToPtr("aBc"))
	assert.Equal(t, "abc", *v)
}

func Test_ToUpper(t *testing.T) {
	assert.Nil(t, ToUpper[string](nil))
	v := ToUpper(gofn.ToPtr("aBc"))
	assert.Equal(t, "ABC", *v)
}

func Test_ToInt64(t *testing.T) {
	assert.Nil(t, ToInt64[int](nil))
	v := ToInt64(gofn.ToPtr(123))
	assert.Equal(t, int64(123), *v)
}

func Test_ToUint64(t *testing.T) {
	assert.Nil(t, ToUint64[int](nil))
	v := ToUint64(gofn.ToPtr(123))
	assert.Equal(t, uint64(123), *v)
}

func Test_ToFloat64(t *testing.T) {
	assert.Nil(t, ToFloat64[int](nil))
	v := ToFloat64(gofn.ToPtr(123.123))
	assert.Equal(t, float64(123.123), *v)
}
