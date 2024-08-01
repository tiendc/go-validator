package stringvalidation

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/tiendc/go-validator/base"
)

var (
	runeLen = utf8.RuneCountInString
)

// RuneLen checks rune length of a string must be in a range
func RuneLen[T base.String](v T, min, max int) (bool, []base.ErrorParam) {
	l := runeLen(string(v))
	return min <= l && l <= max, nil
}

// ByteLen checks byte length of a string must be in a range
func ByteLen[T base.String](v T, min, max int) (bool, []base.ErrorParam) {
	l := len(v)
	return min <= l && l <= max, nil
}

// EQ checks a string must equal to a string
func EQ[T base.String](v T, s T) (bool, []base.ErrorParam) {
	return v == s, nil
}

// In checks a string must be in a list
func In[T base.String](v T, s ...T) (bool, []base.ErrorParam) {
	for i := range s {
		if v == s[i] {
			return true, nil
		}
	}
	return false, nil
}

// NotIn checks a string must be not in a list
func NotIn[T base.String](v T, s ...T) (bool, []base.ErrorParam) {
	for i := range s {
		if v == s[i] {
			return false, nil
		}
	}
	return true, nil
}

// RuneMatch checks a string must be matching a regex on runes
func RuneMatch[T base.String](v T, re *regexp.Regexp) (bool, []base.ErrorParam) {
	return re.MatchReader(strings.NewReader(string(v))), nil
}

// ByteMatch checks a string must be matching a regex on bytes
func ByteMatch[T base.String](v T, re *regexp.Regexp) (bool, []base.ErrorParam) {
	return re.MatchString(string(v)), nil
}
