package validation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/hashicorp/go-multierror"
	"github.com/iancoleman/strcase"
	"github.com/tiendc/gofn"
)

type (
	ErrorParamFormatter interface {
		Format(k string, v any) string
	}

	errorParamFormatter struct {
		k         string
		v         any
		formatter ErrorParamFormatter
	}

	FormatFunc func(reflect.Value) string

	TypedParamFormatter interface {
		ErrorParamFormatter

		SetNumFormatFunc(FormatFunc)
		SetStrFormatFunc(FormatFunc)
		SetBoolFormatFunc(FormatFunc)
		SetSliceFormatFunc(FormatFunc)
		SetMapFormatFunc(FormatFunc)
		SetStructFormatFunc(FormatFunc)
		SetPtrFormatFunc(FormatFunc)
		SetCustomFormatFunc(FormatFunc)
	}

	typedParamFormatter struct {
		numFormatFunc    FormatFunc
		strFormatFunc    FormatFunc
		boolFormatFunc   FormatFunc
		sliceFormatFunc  FormatFunc
		mapFormatFunc    FormatFunc
		structFormatFunc FormatFunc
		ptrFormatFunc    FormatFunc
		customFormatFunc FormatFunc
	}
)

func (f *errorParamFormatter) String() string {
	if f.formatter != nil {
		return f.formatter.Format(f.k, f.v)
	}
	return fmt.Sprintf("%v", f.v)
}

func (f *typedParamFormatter) SetNumFormatFunc(fn FormatFunc) {
	f.numFormatFunc = fn
}
func (f *typedParamFormatter) SetStrFormatFunc(fn FormatFunc) {
	f.strFormatFunc = fn
}
func (f *typedParamFormatter) SetBoolFormatFunc(fn FormatFunc) {
	f.boolFormatFunc = fn
}
func (f *typedParamFormatter) SetSliceFormatFunc(fn FormatFunc) {
	f.sliceFormatFunc = fn
}
func (f *typedParamFormatter) SetMapFormatFunc(fn FormatFunc) {
	f.mapFormatFunc = fn
}
func (f *typedParamFormatter) SetStructFormatFunc(fn FormatFunc) {
	f.structFormatFunc = fn
}
func (f *typedParamFormatter) SetPtrFormatFunc(fn FormatFunc) {
	f.ptrFormatFunc = fn
}
func (f *typedParamFormatter) SetCustomFormatFunc(fn FormatFunc) {
	f.customFormatFunc = fn
}

func (f *typedParamFormatter) Format(_ string, v any) string {
	if f.numFormatFunc == nil && f.strFormatFunc == nil && f.boolFormatFunc == nil &&
		f.sliceFormatFunc == nil && f.mapFormatFunc == nil && f.structFormatFunc == nil &&
		f.ptrFormatFunc == nil && f.customFormatFunc == nil {
		return fmt.Sprintf("%v", v)
	}
	return f.format(reflect.ValueOf(v))
}

func (f *typedParamFormatter) format(v reflect.Value) string {
	if !v.IsValid() {
		return f.customFormat(v)
	}
	// nolint: exhaustive
	switch v.Kind() {
	case reflect.String:
		return f.strFormat(v)
	case reflect.Int, reflect.Uint, reflect.Int64, reflect.Uint64, reflect.Float32, reflect.Float64,
		reflect.Int32, reflect.Uint32, reflect.Int16, reflect.Uint16, reflect.Int8, reflect.Uint8:
		return f.numFormat(v)
	case reflect.Bool:
		return f.boolFormat(v)
	case reflect.Slice, reflect.Array:
		return f.sliceFormat(v)
	case reflect.Map:
		return f.mapFormat(v)
	case reflect.Struct:
		return f.structFormat(v)
	case reflect.Pointer:
		return f.ptrFormat(v)
	default:
		return f.customFormat(v)
	}
}

func (f *typedParamFormatter) strFormat(v reflect.Value) string {
	if f.strFormatFunc == nil {
		return f.customFormat(v)
	}
	return f.strFormatFunc(v)
}

func (f *typedParamFormatter) numFormat(v reflect.Value) string {
	if f.numFormatFunc == nil {
		return f.customFormat(v)
	}
	return f.numFormatFunc(v)
}

func (f *typedParamFormatter) boolFormat(v reflect.Value) string {
	if f.boolFormatFunc == nil {
		return f.customFormat(v)
	}
	return f.boolFormatFunc(v)
}

func (f *typedParamFormatter) sliceFormat(v reflect.Value) string {
	if f.sliceFormatFunc == nil {
		return f.customFormat(v)
	}
	return f.sliceFormatFunc(v)
}

func (f *typedParamFormatter) mapFormat(v reflect.Value) string {
	if f.mapFormatFunc == nil {
		return f.customFormat(v)
	}
	return f.mapFormatFunc(v)
}

func (f *typedParamFormatter) structFormat(v reflect.Value) string {
	if f.structFormatFunc == nil {
		return f.customFormat(v)
	}
	return f.structFormatFunc(v)
}

func (f *typedParamFormatter) ptrFormat(v reflect.Value) string {
	if f.ptrFormatFunc == nil {
		return f.format(v.Elem())
	}
	return f.ptrFormatFunc(v)
}

func (f *typedParamFormatter) customFormat(v reflect.Value) string {
	if f.customFormatFunc == nil {
		if !v.IsValid() {
			return "nil"
		}
		return fmt.Sprintf("%v", v.Interface())
	}
	return f.customFormatFunc(v)
}

func NewTypedParamFormatter() TypedParamFormatter {
	return &typedParamFormatter{}
}

// NewDecimalNumFormatFunc returns a FormatFunc which groups digits of decimal
// For example: '12345' -> '12,345', '12345.6789' -> '12,345.6789'
// To attach this formatter to Error object:
//   - err.TypedParamFormatter().SetNumFormatFunc(NewDecimalNumFormatFunc())
//   - err.TypedParamFormatter().SetNumFormatFunc(NewDecimalNumFormatFunc("%.5f"))
//
// Deprecated: use NewDecimalFormatFunc instead
func NewDecimalNumFormatFunc(floatFmt ...string) FormatFunc {
	return func(v reflect.Value) string {
		var s string
		// nolint: exhaustive
		switch v.Kind() {
		case reflect.Float64, reflect.Float32:
			fmtStr := "%f"
			if len(floatFmt) > 0 {
				fmtStr = floatFmt[len(floatFmt)-1]
			}
			s = fmt.Sprintf(fmtStr, v.Interface())
		default:
			s = fmt.Sprintf("%v", v.Interface())
		}
		return gofn.NumberFmtGroup(s, '.', ',')
	}
}

// NewDecimalFormatFunc returns a FormatFunc which can format and group digits of decimal or integer
// For example: '12345' -> '12,345', '12345.6789' -> '12,345.6789'
// To attach this formatter to Error object:
//   - err.TypedParamFormatter().SetNumFormatFunc(NewDecimalFormatFunc('.', ',', "%.2f"))
func NewDecimalFormatFunc(fractionSep, groupSep byte, floatFmt string) FormatFunc {
	return func(v reflect.Value) string {
		var s string
		// nolint: exhaustive
		switch v.Kind() {
		case reflect.Float64, reflect.Float32:
			fmtStr := floatFmt
			if fmtStr == "" {
				fmtStr = "%f"
			}
			s = fmt.Sprintf(fmtStr, v.Interface())
		default:
			s = fmt.Sprintf("%v", v.Interface())
		}
		return gofn.NumberFmtGroup(s, fractionSep, groupSep)
	}
}

// NewSliceFormatFunc create a new func for formatting a slice
// Sample arguments: leftWrap "[", rightWrap "]", elemSep ", "
func NewSliceFormatFunc(
	elemFormatFunc FormatFunc,
	leftWrap, rightWrap string, elemSep string,
) FormatFunc {
	return func(v reflect.Value) string {
		var sb strings.Builder
		sb.WriteString(leftWrap)
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				sb.WriteString(elemSep)
			}
			sb.WriteString(elemFormatFunc(v.Index(i)))
		}
		sb.WriteString(rightWrap)
		return sb.String()
	}
}

// NewMapFormatFunc create a new func for formatting a map
// Sample arguments: leftWrap "{", rightWrap "}", kvSep ":", elemSep ", "
func NewMapFormatFunc(
	keyFormatFunc, valueFormatFunc FormatFunc,
	leftWrap, rightWrap string, kvSep, entrySep string,
) FormatFunc {
	return func(v reflect.Value) string {
		var sb strings.Builder
		sb.WriteString(leftWrap)
		iter := v.MapRange()
		isFirstItem := true
		for iter.Next() {
			if isFirstItem {
				isFirstItem = false
			} else {
				sb.WriteString(entrySep)
			}
			sb.WriteString(keyFormatFunc(iter.Key()))
			sb.WriteString(kvSep)
			sb.WriteString(valueFormatFunc(iter.Value()))
		}
		sb.WriteString(rightWrap)
		return sb.String()
	}
}

// NewJSONFormatFunc create a format func to format input as JSON output
func NewJSONFormatFunc() FormatFunc {
	return func(v reflect.Value) string {
		s, err := json.Marshal(v.Interface())
		if err != nil {
			panic(err)
		}
		return string(s)
	}
}

// errorBuildDetail builds detail string of error using the error template string
// In case error happens, this function still returns the result string before error happens
func errorBuildDetail(e Error) (detail string, retErr error) {
	detail = e.Template()
	t, err := template.New("error").Parse(detail)
	if err != nil {
		retErr = multierror.Append(retErr, err)
		return
	}

	params, err := errorBuildParams(e, e.ParamFormatter())
	if err != nil {
		retErr = multierror.Append(retErr, err)
	}

	buf := bytes.NewBuffer(make([]byte, 0, 100)) // nolint: gomnd
	err = t.Execute(buf, params)
	if err != nil {
		retErr = multierror.Append(retErr, err)
	} else {
		detail = buf.String()
	}

	return
}

// errorBuildParams builds params of error with inner errors' params handled
func errorBuildParams(e Error, formatter ErrorParamFormatter) (params ErrorParams, err error) {
	params = make(ErrorParams, 10) // nolint: gomnd

	// If there are inner errors, collect all params of them
	for _, inErr := range e.UnwrapAsErrors() {
		prefix := strcase.ToCamel(inErr.Type())
		if prefix != "" {
			prefix += "_"
		}
		pErr := singleErrorBuildParams(inErr, inErr.ParamFormatter(), prefix, params)
		if pErr != nil {
			err = multierror.Append(err, pErr)
		}
	}

	// Build params for the current error
	pErr := singleErrorBuildParams(e, formatter, "", params)
	if pErr != nil {
		err = multierror.Append(err, pErr)
	}

	return params, err
}

// singleErrorBuildParams build params for the specific error
func singleErrorBuildParams(e Error, formatter ErrorParamFormatter, prefix string, params ErrorParams) (err error) {
	key := prefix + "Type"
	params[key] = &errorParamFormatter{k: key, v: e.Type(), formatter: formatter}
	key = prefix + "Value"
	params[key] = &errorParamFormatter{k: key, v: e.Value(), formatter: formatter}
	key = prefix + "ValueType"
	params[key] = &errorParamFormatter{k: key, v: e.ValueType(), formatter: formatter}
	field := e.Field()
	if field != nil {
		key = prefix + "Field"
		params[key] = &errorParamFormatter{k: key, v: field.Name, formatter: formatter}
		key = prefix + "FieldPath"
		params[key] = &errorParamFormatter{k: key, v: field.PathString(false, "."), formatter: formatter}
	} else {
		err = multierror.Append(err, ErrFieldMissing)
		key = prefix + "Field"
		params[key] = &errorParamFormatter{k: key, v: "", formatter: formatter}
		key = prefix + "FieldPath"
		params[key] = &errorParamFormatter{k: key, v: "", formatter: formatter}
	}
	for k, v := range e.Params() {
		key = prefix + k
		params[key] = &errorParamFormatter{k: key, v: v, formatter: formatter}
	}
	return err
}
