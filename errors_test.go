package validation

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_ErrorMod(t *testing.T) {
	err := NewError().SetCustomKey("abc").SetTemplate("template").SetParam("k", "v")

	SetCustomKey("xyz")(err)
	assert.Equal(t, "xyz", err.CustomKey())

	SetTemplate("template {{.Xyz}}")(err)
	assert.Equal(t, "template {{.Xyz}}", err.Template())

	SetParam("k", "vvv")(err)
	assert.Equal(t, "vvv", err.Params()["k"])

	SetField("fieldX", nil)(err)
	assert.Equal(t, "fieldX", err.Field().Name)
	assert.Nil(t, err.Field().Parent)

	SetParamFormatter(nil)(err)
	assert.Nil(t, err.ParamFormatter())
	assert.Nil(t, err.TypedParamFormatter())
	SetParamFormatter(NewTypedParamFormatter())(err)
	assert.NotNil(t, err.ParamFormatter())
	assert.NotNil(t, err.TypedParamFormatter())

	SetNumParamFormatter(func(reflect.Value) string { return "num" })(err)
	SetStrParamFormatter(func(reflect.Value) string { return "str" })(err)
	SetBoolParamFormatter(func(reflect.Value) string { return "bool" })(err)
	SetSliceParamFormatter(func(reflect.Value) string { return "slice" })(err)
	SetMapParamFormatter(func(reflect.Value) string { return "map" })(err)
	SetStructParamFormatter(func(reflect.Value) string { return "struct" })(err)
	SetPtrParamFormatter(func(reflect.Value) string { return "ptr" })(err)
	SetCustomParamFormatter(func(reflect.Value) string { return "custom" })(err)
	assert.Equal(t, "num", err.TypedParamFormatter().Format("k", 123))
	assert.Equal(t, "str", err.TypedParamFormatter().Format("k", "123"))
	assert.Equal(t, "bool", err.TypedParamFormatter().Format("k", true))
	assert.Equal(t, "slice", err.TypedParamFormatter().Format("k", []int{123}))
	assert.Equal(t, "map", err.TypedParamFormatter().Format("k", map[string]any{}))
	assert.Equal(t, "struct", err.TypedParamFormatter().Format("k", struct{}{}))
	assert.Equal(t, "ptr", err.TypedParamFormatter().Format("k", gofn.New(123)))
	assert.Equal(t, "custom", err.TypedParamFormatter().Format("k", func() {}))

	defer func() {
		e := recover()
		assert.Equal(t, "error does not have a TypedParamFormatter attached", e)
	}()
	SetParamFormatter(nil)(err)
	SetNumParamFormatter(func(reflect.Value) string { return "num" })(err)
}

func Test_Field_Path(t *testing.T) {
	field1 := NewField("field1", nil)
	field2 := NewField("field2", field1)
	field3 := NewField("field3", field2)

	assert.Equal(t, "(root).field1.field2.field3", field3.PathString(false, "."))
	assert.Equal(t, "field1.field2.field3", field3.PathString(true, "."))

	assert.Equal(t, "(root)/field1/field2", field2.PathString(false, "/"))
	assert.Equal(t, "field1/field2", field2.PathString(true, "/"))

	assert.Equal(t, "(root).field1", field1.PathString(false, "."))
	assert.Equal(t, "field1", field1.PathString(true, "."))
}

func Test_Error_Impl(t *testing.T) {
	err := NewError()

	_ = err.SetType("type")
	_ = err.SetField(NewField("field", nil))
	_ = err.SetValue("value")
	_ = err.SetValueType("value_type")
	_ = err.SetTemplate("template")
	_ = err.SetParam("k", "v")
	_ = err.SetCustomKey("custom_key")
	assert.Equal(t, "type", err.Type())
	assert.Equal(t, "field", err.Field().Name)
	assert.Equal(t, "value", err.Value())
	assert.Equal(t, "value_type", err.ValueType())
	assert.Equal(t, "template", err.Template())
	assert.Equal(t, "v", err.Params()["k"])
	assert.Equal(t, "custom_key", err.CustomKey())

	detail, e := err.BuildDetail()
	assert.Nil(t, e)
	assert.Equal(t, "template", detail)
	assert.Equal(t, "template", err.Error())
	assert.Equal(t, "template", err.String())
	assert.Equal(t, []error{}, err.Unwrap())
	var errs Errors
	assert.Equal(t, errs, err.UnwrapAsErrors())
}

func Test_Errors(t *testing.T) {
	errs := Errors{
		NewError().SetTemplate("template_1"),
		NewError().SetTemplate("template_2"),
	}
	assert.Equal(t, "template_1\ntemplate_2", errs.Error())
}
