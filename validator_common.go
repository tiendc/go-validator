package validation

import (
	commonFunc "github.com/tiendc/go-validator/base/common"
)

// Nil validates the input pointer to be `nil`
func Nil[T any](v *T) SingleValidator {
	return call1[*T]("nil", "", commonFunc.Nil[T])(v)
}

// NotNil validates the input pointer to be not `nil`
func NotNil[T any](v *T) SingleValidator {
	return call1[*T]("not_nil", "", commonFunc.NotNil[T])(v)
}

// Required validates the input to be required.
// Required value must be not:
//   - zero value (0, "", nil, false)
//   - empty slice, array, map, channel
//   - pointer points to zero value
func Required(v any) SingleValidator {
	return call1[any]("required", "", commonFunc.Required)(v)
}
