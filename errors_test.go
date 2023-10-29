package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
