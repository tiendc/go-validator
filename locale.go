package validation

import (
	enTemplates "github.com/tiendc/go-validator/internal/templates/en"
)

const (
	langEn = "en"
)

// TemplateProvider interface for providing template for generating error messages
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

// DefaultLang default language used in the template
func DefaultLang() string {
	return defaultLang
}

// SetDefaultLang set default language used in the template
func SetDefaultLang(lang string) {
	defaultLang = lang
}

// GetTemplateProvider gets current template provider
func GetTemplateProvider(lang string) TemplateProvider {
	return fmtTemplates[lang]
}

// SetTemplateProvider sets current template provider
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
