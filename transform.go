package validation

import (
	"strings"

	"github.com/tiendc/go-validator/base"
)

// ToLower transforms characters of a string to lowercase
// E.g. validation.StrIsEmail(validation.ToLower(&req.Email))
func ToLower[T base.String](v *T) *T {
	if v == nil {
		return v
	}
	vv := T(strings.ToLower(string(*v)))
	return &vv
}

func ToUpper[T base.String](v *T) *T {
	if v == nil {
		return v
	}
	vv := T(strings.ToUpper(string(*v)))
	return &vv
}

func ToInt64[T base.Number](v *T) *int64 {
	if v == nil {
		return nil
	}
	vv := int64(*v)
	return &vv
}

func ToUint64[T base.Number](v *T) *uint64 {
	if v == nil {
		return nil
	}
	vv := uint64(*v)
	return &vv
}

func ToFloat64[T base.Number](v *T) *float64 {
	if v == nil {
		return nil
	}
	vv := float64(*v)
	return &vv
}
