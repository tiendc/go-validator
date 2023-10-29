package validation

import (
	"errors"
	"strings"

	"github.com/tiendc/gofn"
)

var (
	ErrTypeUnsupported = errors.New("type unsupported")
	ErrFieldMissing    = errors.New("field missing")
)

const (
	rootField = "(root)"
)

type (
	// ErrorParams is a map of params specific to each error
	ErrorParams map[string]any

	Field struct {
		Name   string
		Parent *Field
	}

	// Error is the interface for all validation errors in this lib
	Error interface {
		Type() string
		SetType(string) Error

		Field() *Field
		SetField(*Field) Error

		Value() any
		SetValue(any) Error

		ValueType() string
		SetValueType(string) Error

		Template() string
		SetTemplate(string) Error

		Params() ErrorParams
		SetParam(k string, v any) Error

		// ParamFormatter formatter is used to format the error params
		// By default it is TypedParamFormatter
		ParamFormatter() ErrorParamFormatter
		// TypedParamFormatter get TypedParamFormatter attached to the error
		// This will return nil when the attached formatter is not a TypedParamFormatter
		TypedParamFormatter() TypedParamFormatter
		SetParamFormatter(ErrorParamFormatter) Error

		CustomKey() any
		SetCustomKey(any) Error

		BuildDetail() (string, error)
		ParamsWithFormatter() ErrorParams

		// String implements fmt.Stringer interface
		// This function calls BuildDetail() without raising error
		// Should use BuildDetail() for more controls on error
		String() string
		// Error implement error interface
		// See String() string
		Error() string

		// Unwrap implements errors.Unwrap
		Unwrap() []error
		UnwrapAsErrors() Errors
	}

	Errors []Error

	ErrorMod func(Error)

	// errorImpl implementation of Error type
	// nolint: errname
	errorImpl struct {
		errorType       string
		field           *Field
		value           any
		valueType       string
		template        string
		params          ErrorParams
		paramsFormatter ErrorParamFormatter
		customKey       any

		wrappedErrors []Error
	}
)

func NewField(name string, parent *Field) *Field {
	return &Field{name, parent}
}

func (c *Field) Path(skipRoot bool) []string {
	path := []string{c.Name}
	t := c.Parent
	for t != nil {
		path = append(path, t.Name)
		t = t.Parent
	}
	if !skipRoot {
		path = append(path, rootField)
	}
	return gofn.Reverse(path)
}

func (c *Field) PathString(skipRoot bool, sep string) string {
	return strings.Join(c.Path(skipRoot), sep)
}

func NewError() Error {
	return &errorImpl{
		params:          ErrorParams{},
		paramsFormatter: NewTypedParamFormatter(),
	}
}

func (e *errorImpl) Type() string {
	return e.errorType
}

func (e *errorImpl) SetType(errorType string) Error {
	e.errorType = errorType
	return e
}

func (e *errorImpl) Field() *Field {
	return e.field
}

func (e *errorImpl) SetField(field *Field) Error {
	e.field = field
	return e
}

func (e *errorImpl) Value() any {
	return e.value
}

func (e *errorImpl) SetValue(value any) Error {
	e.value = value
	return e
}

func (e *errorImpl) ValueType() string {
	return e.valueType
}

func (e *errorImpl) SetValueType(valueType string) Error {
	e.valueType = valueType
	return e
}

func (e *errorImpl) Template() string {
	return e.template
}

func (e *errorImpl) SetTemplate(template string) Error {
	e.template = template
	return e
}

func (e *errorImpl) Params() ErrorParams {
	return e.params
}

func (e *errorImpl) SetParam(key string, val any) Error {
	if e.params == nil {
		e.params = ErrorParams{}
	}
	e.params[key] = val
	return e
}

func (e *errorImpl) ParamFormatter() ErrorParamFormatter {
	return e.paramsFormatter
}

func (e *errorImpl) TypedParamFormatter() TypedParamFormatter {
	return e.paramsFormatter.(TypedParamFormatter) // nolint: forcetypeassert
}

func (e *errorImpl) SetParamFormatter(formatter ErrorParamFormatter) Error {
	e.paramsFormatter = formatter
	return e
}

func (e *errorImpl) CustomKey() any {
	return e.customKey
}

func (e *errorImpl) SetCustomKey(key any) Error {
	e.customKey = key
	return e
}

func (e *errorImpl) BuildDetail() (string, error) {
	return errorBuildDetail(e)
}

func (e *errorImpl) ParamsWithFormatter() ErrorParams {
	params, _ := errorBuildParams(e, e.paramsFormatter)
	return params
}

func (e *errorImpl) String() string {
	str, _ := errorBuildDetail(e)
	return str
}

func (e *errorImpl) Error() string {
	return e.String()
}

func (e *errorImpl) Unwrap() []error {
	errs := make([]error, 0, len(e.wrappedErrors))
	for _, err := range e.wrappedErrors {
		errs = append(errs, err)
	}
	return errs
}

func (e *errorImpl) UnwrapAsErrors() Errors {
	return e.wrappedErrors
}

func (e Errors) Error() string {
	if len(e) == 0 {
		return ""
	}
	errStr := ""
	for i, err := range e {
		if i > 0 {
			errStr += "\n"
		}
		errStr += err.Error()
	}
	return errStr
}

func SetField(name string, parent *Field) ErrorMod {
	return func(err Error) {
		_ = err.SetField(NewField(name, parent))
	}
}

func SetCustomKey(key any) ErrorMod {
	return func(err Error) {
		_ = err.SetCustomKey(key)
	}
}

func SetTemplate(template string) ErrorMod {
	return func(err Error) {
		_ = err.SetTemplate(template)
	}
}

func SetParam(key string, val any) ErrorMod {
	return func(err Error) {
		_ = err.SetParam(key, val)
	}
}

func SetParamFormatter(formatter ErrorParamFormatter) ErrorMod {
	return func(err Error) {
		_ = err.SetParamFormatter(formatter)
	}
}
