package config

import (
	"log"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Env struct {
	SecretKey     	string      `env:"SECRET_KEY,required"`
	DSN         	 	string      `env:"DSN,required"`
	ExpiredToken  	int         `env:"EXPIRED_TOKEN,required" envDefault:"240"`
	Env             string			`env:"ENV,required" envDefault:"development"`
}

var cfg = Env{}

func ValidateEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	err = env.Parse(&cfg)
  if err != nil {
		log.Fatalf("%+v\n", err)
  }
}

func GetEnv() Env {
	return cfg
}