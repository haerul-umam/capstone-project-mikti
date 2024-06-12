package helper

import (
	"github.com/go-playground/validator/v10"
)

func NewValidator() *Validator {
	validator := validator.New()
	validator.RegisterValidation("status_check", StatusValidate)
	return &Validator{validator}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func StatusValidate(fi validator.FieldLevel) bool {
	if fi.Field().String() == "MENUNGGU" || fi.Field().String() == "DITERIMA" || fi.Field().String() == "DITOLAK" {
		return true
	}

	return false
}
