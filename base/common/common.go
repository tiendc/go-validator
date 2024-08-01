package commonvalidation

import (
	"github.com/tiendc/gofn"

	"github.com/tiendc/go-validator/base"
)

// Nil checks input must be `nil`
func Nil[T any](v *T) (bool, []base.ErrorParam) {
	if v == nil {
		return true, nil
	}
	return false, nil
}

// NotNil checks input must be not `nil`
func NotNil[T any](v *T) (bool, []base.ErrorParam) {
	if v != nil {
		return true, nil
	}
	return false, nil
}

// Required checks input must be a valid value.
// A required value must be not:
//   - zero value (0, "", nil, false)
//   - empty slice, array, map, channel
//   - pointer points to zero value
func Required(v any) (bool, []base.ErrorParam) {
	if v == nil || gofn.FirstTrue(nil, v) == nil {
		return false, nil
	}
	return true, nil
}
