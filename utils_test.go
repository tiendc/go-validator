package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
