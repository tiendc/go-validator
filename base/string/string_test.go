package stringvalidation

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_RuneLen_ByteLen(t *testing.T) {
	assert.True(t, gofn.Head(RuneLen("chào", 1, 4)))
	assert.True(t, gofn.Head(RuneLen("abc", 1, 3)))
	assert.False(t, gofn.Head(RuneLen("chào", 1, 3)))
	assert.False(t, gofn.Head(RuneLen("abc", 1, 2)))

	assert.True(t, gofn.Head(ByteLen("abc", 1, 5)))
	assert.True(t, gofn.Head(ByteLen("chào", 1, 5)))
	assert.False(t, gofn.Head(ByteLen("abc", 1, 2)))
	assert.False(t, gofn.Head(ByteLen("chào", 1, 4)))
}

func Test_EQ(t *testing.T) {
	assert.True(t, gofn.Head(EQ("", "")))
	assert.True(t, gofn.Head(EQ("abc", "abc")))
	assert.True(t, gofn.Head(EQ("chào", "chào")))
	assert.False(t, gofn.Head(EQ("abc", "aBc")))
	assert.False(t, gofn.Head(EQ("abc", "")))
	assert.False(t, gofn.Head(EQ("chào", "chao")))
}

func Test_In(t *testing.T) {
	assert.True(t, gofn.Head(In("abc", "abc")))
	assert.True(t, gofn.Head(In("abc", "", "abc", "123")))
	assert.True(t, gofn.Head(In("", "aBc", "")))
	assert.False(t, gofn.Head(In("abc", "aBc", "123", "")))
	assert.False(t, gofn.Head(In("", "aBc", "123")))
}

func Test_NotIn(t *testing.T) {
	assert.True(t, gofn.Head(NotIn("abc", "abC")))
	assert.True(t, gofn.Head(NotIn("", "aBc", "123")))
	assert.False(t, gofn.Head(NotIn("abc", "abc", "123", "")))
	assert.False(t, gofn.Head(NotIn("", "aBc", "123", "")))
}

func Test_Match(t *testing.T) {
	re := regexp.MustCompile("[0-9]+")
	assert.True(t, gofn.Head(RuneMatch("1234", re)))
	assert.False(t, gofn.Head(RuneMatch("abc", re)))
	assert.True(t, gofn.Head(ByteMatch("1234", re)))
	assert.False(t, gofn.Head(ByteMatch("abc", re)))
}
