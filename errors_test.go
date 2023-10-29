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
