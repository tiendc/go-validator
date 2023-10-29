package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultLang(t *testing.T) {
	assert.Equal(t, langEn, DefaultLang())
	SetDefaultLang("vi")
	assert.Equal(t, "vi", DefaultLang())
}

func Test_GetTemplateProvider(t *testing.T) {
	SetDefaultLang(langEn)
	assert.NotNil(t, GetTemplateProvider(langEn))
	assert.Nil(t, GetTemplateProvider("vi"))
}

func Test_SetTemplateProvider(t *testing.T) {
	provider := &templateProvider{}
	SetTemplateProvider("vi", provider)
	assert.Equal(t, provider, GetTemplateProvider("vi"))
}

func Test_getFmtTemplate(t *testing.T) {
	err := NewError()
	_ = err.SetType("required")

	SetDefaultLang(langEn)
	assert.Equal(t, `{{.Field}} is required`, getFmtTemplate(err))
	SetDefaultLang("xy")
	assert.Equal(t, `{{.Field}} is required`, getFmtTemplate(err))
}
