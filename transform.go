package validation

import (
	"strings"

	"github.com/tiendc/go-validator/base"
)

// ToLower transforms characters of a string to lowercase
//
// Example: validation.StrIsEmail(validation.ToLower(&req.Email))
func ToLower[T base.String](v *T) *T {
	if v == nil {
		return v
	}
	vv := T(strings.ToLower(string(*v)))
	return &vv
}

// ToUpper transforms characters of a string to uppercase
func ToUpper[T base.String](v *T) *T {
	if v == nil {
		return v
	}
	vv := T(strings.ToUpper(string(*v)))
	return &vv
}

// ToInt64 transforms a number to int64 value
func ToInt64[T base.Number](v *T) *int64 {
	if v == nil {
		return nil
	}
	vv := int64(*v)
	return &vv
}

// ToUint64 transforms a number to uint64 value
func ToUint64[T base.Number](v *T) *uint64 {
	if v == nil {
		return nil
	}
	vv := uint64(*v)
	return &vv
}

// ToFloat64 transforms a number to float64 value
func ToFloat64[T base.Number](v *T) *float64 {
	if v == nil {
		return nil
	}
	vv := float64(*v)
	return &vv
}
