package validation

import (
	enTemplates "github.com/tiendc/go-validator/internal/templates/en"
)

const (
	langEn = "en"
)

type TemplateProvider interface {
	Get(Error) string
}

type templateProvider struct {
	fmtTemplates map[string]string
}

func (t *templateProvider) Get(e Error) string {
	key := e.ValueType()
	if key != "" {
		key += "_"
	}
	key += e.Type()
	return t.fmtTemplates[key]
}

var (
	defaultLang  = langEn
	fmtTemplates = map[string]TemplateProvider{
		langEn: &templateProvider{fmtTemplates: enTemplates.FmtTemplates},
	}
)

func DefaultLang() string {
	return defaultLang
}

func SetDefaultLang(lang string) {
	defaultLang = lang
}

func GetTemplateProvider(lang string) TemplateProvider {
	return fmtTemplates[lang]
}

func SetTemplateProvider(lang string, provider TemplateProvider) {
	fmtTemplates[lang] = provider
}

func getFmtTemplate(e Error) string {
	provider := GetTemplateProvider(defaultLang)
	if provider == nil {
		provider = GetTemplateProvider(langEn)
	}
	return provider.Get(e)
}
