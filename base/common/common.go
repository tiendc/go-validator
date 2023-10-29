package commonvalidation

import (
	"github.com/tiendc/gofn"

	"github.com/tiendc/go-validator/base"
)

func Nil[T any](v *T) (bool, []base.ErrorParam) {
	if v == nil {
		return true, nil
	}
	return false, nil
}

func NotNil[T any](v *T) (bool, []base.ErrorParam) {
	if v != nil {
		return true, nil
	}
	return false, nil
}

func Required(v any) (bool, []base.ErrorParam) {
	if v == nil || gofn.FirstTrue(nil, v) == nil {
		return false, nil
	}
	return true, nil
}
