package config

import (
	val "rohmatext/ore-note/internal/delivery/http/validator"

	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/en"
)

func NewValidator() *val.Validator {
	v := validator.New()
	locale := en.New()
	uni := ut.New(locale, locale)
	trans, _ := uni.GetTranslator("en")
	translations.RegisterDefaultTranslations(v, trans)
	return &val.Validator{
		Validator:  v,
		Translator: trans,
	}
}
