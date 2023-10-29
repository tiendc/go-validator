package validation

import (
	"github.com/tiendc/go-validator/base"
)

// errorBuild builds Error object based on the given params
func errorBuild(typ string, vTyp string, value any, wrapErrs Errors, params ...base.ErrorParam) Error {
	err := &errorImpl{
		errorType:       typ,
		valueType:       vTyp,
		value:           value,
		params:          ErrorParams{},
		paramsFormatter: NewTypedParamFormatter(),
		wrappedErrors:   wrapErrs,
	}
	for _, p := range params {
		err.params[p.Key] = p.Value
	}
	err.template = getFmtTemplate(err)
	return err
}
