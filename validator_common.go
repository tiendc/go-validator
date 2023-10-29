package validation

import (
	commonFunc "github.com/tiendc/go-validator/base/common"
)

func Nil[T any](v *T) SingleValidator {
	return call1[*T]("nil", "", commonFunc.Nil[T])(v)
}

func NotNil[T any](v *T) SingleValidator {
	return call1[*T]("not_nil", "", commonFunc.NotNil[T])(v)
}

func Required(v any) SingleValidator {
	return call1[any]("required", "", commonFunc.Required)(v)
}
