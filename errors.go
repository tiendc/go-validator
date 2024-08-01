package validation

import (
	"errors"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/tiendc/gofn"
)

// Errors can be returned from the lib
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

	// Field represents a validating field
	Field struct {
		Name   string
		Parent *Field
	}

	// Error is the interface for all validation errors in this lib
	Error interface {
		// Type gets type of error
		Type() string
		// SetType sets type of error
		SetType(string) Error

		// Field gets validating field
		Field() *Field
		// SetField sets validating field
		SetField(*Field) Error

		// Value gets validating value
		Value() any
		// SetValue gets validating value
		SetValue(any) Error

		// ValueType gets type of validating value
		ValueType() string
		// SetValueType sets type of validating value
		SetValueType(string) Error

		// Template gets template used to generating error message
		Template() string
		// SetTemplate sets template of error
		SetTemplate(string) Error

		// Params gets params of error
		Params() ErrorParams
		// SetParam sets params of error
		SetParam(k string, v any) Error

		// ParamFormatter formatter is used to format the error params
		// By default it is TypedParamFormatter
		ParamFormatter() ErrorParamFormatter
		// TypedParamFormatter get TypedParamFormatter attached to the error
		// This will return nil when the attached formatter is not a TypedParamFormatter
		TypedParamFormatter() TypedParamFormatter
		// SetParamFormatter sets params formatter of error
		SetParamFormatter(ErrorParamFormatter) Error

		// CustomKey gets custom key of error
		CustomKey() any
		// SetCustomKey sets custom key of error
		SetCustomKey(any) Error

		// BuildDetail builds error detailed message
		BuildDetail() (string, error)
		// ParamsWithFormatter gets params with wrapping by the formatter of the error
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
		// UnwrapAsErrors unwraps the error as `Errors` type
		UnwrapAsErrors() Errors
	}

	// Errors slice type for `Error` objects
	Errors []Error

	// ErrorMod function used to modify an `Error` object
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

// NewField creates a new Field object
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

// NewError creates a new Error object
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
	params := gofn.MapUpdate(make(ErrorParams, len(e.params)), e.params)
	// Collect all inner errors' params
	for _, inErr := range e.wrappedErrors {
		prefix := strcase.ToCamel(inErr.Type())
		if prefix != "" {
			prefix += "_"
		}
		for k, v := range inErr.Params() {
			params[prefix+k] = v
		}
	}
	return params
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
	if e.paramsFormatter == nil {
		return nil
	}
	typedFmt, _ := e.paramsFormatter.(TypedParamFormatter)
	return typedFmt
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
	var sb strings.Builder
	for i, err := range e {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(err.Error())
	}
	return sb.String()
}

// SetField returns a ErrorMod function to set field of error
func SetField(name string, parent *Field) ErrorMod {
	return func(err Error) {
		_ = err.SetField(NewField(name, parent))
	}
}

// SetCustomKey returns a ErrorMod function to set custom key of error
func SetCustomKey(key any) ErrorMod {
	return func(err Error) {
		_ = err.SetCustomKey(key)
	}
}

// SetTemplate returns a ErrorMod function to set template of error
func SetTemplate(template string) ErrorMod {
	return func(err Error) {
		_ = err.SetTemplate(template)
	}
}

// SetParam returns a ErrorMod function to set a param of error
func SetParam(key string, val any) ErrorMod {
	return func(err Error) {
		_ = err.SetParam(key, val)
	}
}

// SetParamFormatter returns a ErrorMod function to set params formatter of error
func SetParamFormatter(formatter ErrorParamFormatter) ErrorMod {
	return func(err Error) {
		_ = err.SetParamFormatter(formatter)
	}
}

// SetNumParamFormatter returns a ErrorMod function to set format function for numbers
func SetNumParamFormatter(formatFunc FormatFunc) ErrorMod {
	return func(err Error) {
		getTypedParamFormatterOrPanic(err).SetNumFormatFunc(formatFunc)
	}
}

// SetStrParamFormatter returns a ErrorMod function to set format function for strings
func SetStrParamFormatter(formatFunc FormatFunc) ErrorMod {
	return func(err Error) {
		getTypedParamFormatterOrPanic(err).SetStrFormatFunc(formatFunc)
	}
}

// SetBoolParamFormatter returns a ErrorMod function to set format function for bools
func SetBoolParamFormatter(formatFunc FormatFunc) ErrorMod {
	return func(err Error) {
		getTypedParamFormatterOrPanic(err).SetBoolFormatFunc(formatFunc)
	}
}

// SetSliceParamFormatter returns a ErrorMod function to set format function for slices
func SetSliceParamFormatter(formatFunc FormatFunc) ErrorMod {
	return func(err Error) {
		getTypedParamFormatterOrPanic(err).SetSliceFormatFunc(formatFunc)
	}
}

// SetMapParamFormatter returns a ErrorMod function to set format function for maps
func SetMapParamFormatter(formatFunc FormatFunc) ErrorMod {
	return func(err Error) {
		getTypedParamFormatterOrPanic(err).SetMapFormatFunc(formatFunc)
	}
}

// SetStructParamFormatter returns a ErrorMod function to set format function for structs
func SetStructParamFormatter(formatFunc FormatFunc) ErrorMod {
	return func(err Error) {
		getTypedParamFormatterOrPanic(err).SetStructFormatFunc(formatFunc)
	}
}

// SetPtrParamFormatter returns a ErrorMod function to set format function for pointers
func SetPtrParamFormatter(formatFunc FormatFunc) ErrorMod {
	return func(err Error) {
		getTypedParamFormatterOrPanic(err).SetPtrFormatFunc(formatFunc)
	}
}

// SetCustomParamFormatter returns a ErrorMod function to set custom format function
func SetCustomParamFormatter(formatFunc FormatFunc) ErrorMod {
	return func(err Error) {
		getTypedParamFormatterOrPanic(err).SetCustomFormatFunc(formatFunc)
	}
}

// getTypedParamFormatterOrPanic returns the TypedParamFormatter associated with the error or panic if unset
func getTypedParamFormatterOrPanic(err Error) TypedParamFormatter {
	formatter := err.TypedParamFormatter()
	if formatter == nil {
		panic("error does not have a TypedParamFormatter attached")
	}
	return formatter
}
