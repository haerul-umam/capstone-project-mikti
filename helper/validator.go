package helper

import (
	"log"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
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

type config struct {
	SecretKey     	string      `env:"SECRET_KEY,required"`
	DSN         	 	string      `env:"DSN,required"`
	ExpiredToken  	int         `env:"EXPIRED_TOKEN,required" envDefault:"240"`
}

func ValidateEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}
	_, err = env.ParseAs[config]()
  if err != nil {
		log.Fatalf("%+v\n", err)
  }
}