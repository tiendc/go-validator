package numbervalidation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_EQ(t *testing.T) {
	assert.True(t, gofn.Head(EQ(123, 123)))
	assert.True(t, gofn.Head(EQ(-1.1, -1.1)))
	assert.False(t, gofn.Head(EQ(100, 101)))
	assert.False(t, gofn.Head(EQ(1.123, 1.1234)))
}

func Test_GT(t *testing.T) {
	assert.True(t, gofn.Head(GT(123, 122)))
	assert.True(t, gofn.Head(GT(0, -1)))
	assert.False(t, gofn.Head(GT(100, 100)))
	assert.True(t, gofn.Head(GTE(0, -1)))
	assert.True(t, gofn.Head(GTE(100, 100)))
	assert.True(t, gofn.Head(GTE(-1, -1)))
}

func Test_LT(t *testing.T) {
	assert.True(t, gofn.Head(LT(123, 124)))
	assert.True(t, gofn.Head(LT(-1, 0)))
	assert.False(t, gofn.Head(LT(100, 100)))
	assert.True(t, gofn.Head(LTE(-1, -1)))
	assert.True(t, gofn.Head(LTE(100, 100)))
	assert.True(t, gofn.Head(LTE(-1, 0)))
}

func Test_Range(t *testing.T) {
	assert.True(t, gofn.Head(Range(200, 100, 200)))
	assert.True(t, gofn.Head(Range(100, 100, 200)))
	assert.True(t, gofn.Head(Range(150, 100, 200)))
	assert.False(t, gofn.Head(Range(99, 100, 200)))
}

func Test_In(t *testing.T) {
	assert.True(t, gofn.Head(In(1, 1)))
	assert.True(t, gofn.Head(In(1, 0, 1, 2)))
	assert.False(t, gofn.Head(In(0, 1, 2)))
}

func Test_NotIn(t *testing.T) {
	assert.True(t, gofn.Head(NotIn(1, 0)))
	assert.True(t, gofn.Head(NotIn(1, 2, 3)))
	assert.False(t, gofn.Head(NotIn(0, 1, 0, 2)))
}

func Test_DivisibleBy(t *testing.T) {
	assert.True(t, gofn.Head(DivisibleBy(10, 5)))
	assert.True(t, gofn.Head(DivisibleBy(22, 11)))
	assert.True(t, gofn.Head(DivisibleBy(10, 1)))
	assert.True(t, gofn.Head(DivisibleBy(0, 123)))
	assert.False(t, gofn.Head(DivisibleBy(8, 3)))
	assert.False(t, gofn.Head(DivisibleBy(8, 0)))
}

func Test_JsSafeInt(t *testing.T) {
	assert.True(t, gofn.Head(JsSafeInt(10)))
	assert.True(t, gofn.Head(JsSafeInt(9007199254740991)))
	assert.True(t, gofn.Head(JsSafeInt(-9007199254740991)))

	assert.False(t, gofn.Head(JsSafeInt(9007199254740992)))
	assert.False(t, gofn.Head(JsSafeInt(-9007199254740992)))
}
