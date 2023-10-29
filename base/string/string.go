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

func RuneLen[T base.String](v T, min, max int) (bool, []base.ErrorParam) {
	l := runeLen(string(v))
	return min <= l && l <= max, nil
}

func ByteLen[T base.String](v T, min, max int) (bool, []base.ErrorParam) {
	l := len(v)
	return min <= l && l <= max, nil
}

func EQ[T base.String](v T, s T) (bool, []base.ErrorParam) {
	return v == s, nil
}

func In[T base.String](v T, s ...T) (bool, []base.ErrorParam) {
	for i := range s {
		if v == s[i] {
			return true, nil
		}
	}
	return false, nil
}

func NotIn[T base.String](v T, s ...T) (bool, []base.ErrorParam) {
	for i := range s {
		if v == s[i] {
			return false, nil
		}
	}
	return true, nil
}

func RuneMatch[T base.String](v T, re *regexp.Regexp) (bool, []base.ErrorParam) {
	return re.MatchReader(strings.NewReader(string(v))), nil
}

func ByteMatch[T base.String](v T, re *regexp.Regexp) (bool, []base.ErrorParam) {
	return re.MatchString(string(v)), nil
}
