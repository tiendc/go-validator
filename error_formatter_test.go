package validation

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_TypedParamFormatter(t *testing.T) {
	formatter := NewTypedParamFormatter()
	assert.Equal(t, "123", formatter.Format("", 123))
	assert.Equal(t, "abc", formatter.Format("", "abc"))
	assert.Equal(t, "true", formatter.Format("", true))
	assert.True(t, strings.HasPrefix(formatter.Format("", gofn.New("abc")), "0x"))
	assert.Equal(t, "[abc 123]", formatter.Format("", []any{"abc", 123}))
	assert.Equal(t, "map[1:abc 2:123]", formatter.Format("", map[int]any{1: "abc", 2: 123}))

	type S struct {
		I int
		U uint
	}
	assert.Equal(t, "{111 222}", formatter.Format("", S{I: 111, U: 222}))

	jsonFormatFunc := NewJSONFormatFunc()
	// Use custom format function only
	formatter.SetCustomFormatFunc(jsonFormatFunc)

	assert.Equal(t, "123", formatter.Format("", 123))
	assert.Equal(t, `"abc"`, formatter.Format("", "abc"))
	assert.Equal(t, "true", formatter.Format("", true))
	assert.Equal(t, `"abc"`, formatter.Format("", gofn.New("abc")))
	assert.Equal(t, `["abc",123]`, formatter.Format("", []any{"abc", 123}))
	v := formatter.Format("", map[int]any{1: "abc", 2: 123})
	assert.True(t, v == `{"1":"abc","2":123}` || v == `{"2":123,"1":"abc"}`)

	formatter.SetNumFormatFunc(jsonFormatFunc)
	formatter.SetStrFormatFunc(jsonFormatFunc)
	formatter.SetBoolFormatFunc(jsonFormatFunc)
	formatter.SetSliceFormatFunc(jsonFormatFunc)
	formatter.SetMapFormatFunc(jsonFormatFunc)
	formatter.SetStructFormatFunc(jsonFormatFunc)
	formatter.SetPtrFormatFunc(jsonFormatFunc)
	formatter.SetCustomFormatFunc(jsonFormatFunc)

	assert.Equal(t, "123", formatter.Format("", 123))
	assert.Equal(t, `"abc"`, formatter.Format("", "abc"))
	assert.Equal(t, "true", formatter.Format("", true))
	assert.Equal(t, `"abc"`, formatter.Format("", gofn.New("abc")))
	assert.Equal(t, `["abc",123]`, formatter.Format("", []any{"abc", 123}))
	v = formatter.Format("", map[int]any{1: "abc", 2: 123})
	assert.True(t, v == `{"1":"abc","2":123}` || v == `{"2":123,"1":"abc"}`)
}

func Test_NewJSONFormatFunc(t *testing.T) {
	fmtFunc := NewJSONFormatFunc()
	assert.Equal(t, `"abc"`, fmtFunc(reflect.ValueOf("abc")))
	assert.Equal(t, "123", fmtFunc(reflect.ValueOf(123)))
	assert.Equal(t, `["abc",123]`, fmtFunc(reflect.ValueOf([]any{"abc", 123})))
	assert.Equal(t, `{"1":"abc","2":123}`, fmtFunc(reflect.ValueOf(map[int]any{1: "abc", 2: 123})))
}

func Test_NewSliceFormatFunc(t *testing.T) {
	itemFmtFunc := NewJSONFormatFunc()
	fmtFunc := NewSliceFormatFunc(itemFmtFunc, "[", "]", ",")
	assert.Equal(t, `["abc",123]`, fmtFunc(reflect.ValueOf([]any{"abc", 123})))
}

func Test_NewMapFormatFunc(t *testing.T) {
	kvFmtFunc := NewJSONFormatFunc()
	fmtFunc := NewMapFormatFunc(kvFmtFunc, kvFmtFunc, "{", "}", ":", ",")
	v := fmtFunc(reflect.ValueOf(map[int]any{1: "abc", 2: 123}))
	assert.True(t, v == `{1:"abc",2:123}` || v == `{2:123,1:"abc"}`)
}

func Test_NewDecimalNumFormatFunc(t *testing.T) {
	fmtFunc := NewDecimalNumFormatFunc()
	assert.Equal(t, "12,345", fmtFunc(reflect.ValueOf(12345)))
	assert.Equal(t, "1,234,567.123457", fmtFunc(reflect.ValueOf(1234567.1234567)))
	fmtFunc = NewDecimalNumFormatFunc("%.5f")
	assert.Equal(t, "12,345", fmtFunc(reflect.ValueOf(12345)))
	assert.Equal(t, "1,234,567.12346", fmtFunc(reflect.ValueOf(1234567.1234567)))
}

func Test_errorBuildDetail(t *testing.T) {
	err := NewError().
		SetCustomKey("customKey").
		SetTemplate("'{{.Value}}': {{.Field}} is invalid").
		SetParam("k", "v").
		SetValue("value2").
		SetField(NewField("field2", NewField("field1", nil)))

	detail, buildErr := errorBuildDetail(err)
	assert.Nil(t, buildErr)
	assert.Equal(t, "'value2': field2 is invalid", detail)
}
