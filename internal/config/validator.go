package config

import (
	"fmt"
	"reflect"
	validate "rohmatext/ore-note/internal/delivery/http/validator"
	"strings"

	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/en"
	"gorm.io/gorm"
)

func NewValidator(db *gorm.DB) *validate.Validator {
	v := validator.New()
	locale := en.New()
	uni := ut.New(locale, locale)
	trans, _ := uni.GetTranslator("en")
	translations.RegisterDefaultTranslations(v, trans)
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("json")
		if name == "-" {
			return ""
		}
		return name
	})

	v.RegisterValidation("unique_table", func(fl validator.FieldLevel) bool {
		raw := fl.Param()
		parts := strings.Split(raw, ".")

		if len(parts) != 2 {
			return false
		}

		table := parts[0]
		column := parts[1]

		value := fl.Field().String()

		parent := fl.Parent()
		if parent.Kind() == reflect.Pointer {
			parent = parent.Elem()
		}
		var excludeID uint = 0

		idField := parent.FieldByName("ID")
		if idField.IsValid() && idField.Kind() == reflect.Uint {
			excludeID = uint(idField.Uint())
		}

		query := db.Table(table).Where(fmt.Sprintf("%s = ?", column), value)

		if excludeID > 0 {
			query = query.Where("id != ?", excludeID)
		}

		var count int64
		if err := query.Count(&count).Error; err != nil {
			return false
		}

		return count == 0
	})

	v.RegisterValidation("exists", func(fl validator.FieldLevel) bool {
		raw := fl.Param()
		parts := strings.Split(raw, ".")
		if len(parts) != 2 {
			return false
		}

		table := parts[0]
		column := parts[1]

		var value any
		switch fl.Field().Kind() {
		case reflect.String:
			value = fl.Field().String()
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			value = fl.Field().Int()
		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = fl.Field().Uint()
		default:
			return false
		}

		var count int64
		err := db.Table(table).
			Where(fmt.Sprintf("%s = ?", column), value).
			Count(&count).Error

		if err != nil {
			return false
		}

		return count > 0
	})

	v.RegisterTranslation("unique_table", trans, func(ut ut.Translator) error {
		return ut.Add("unique_table", "{0} has already been taken.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("unique_table", fe.Field())
		return t
	})

	v.RegisterTranslation("exists", trans, func(ut ut.Translator) error {
		return ut.Add("exists", "{0} is not invalid.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("exists", fe.Field())
		return t
	})

	return &validate.Validator{
		Validator:  v,
		Translator: trans,
	}
}
