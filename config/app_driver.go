package config

import (
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
)

type App struct {
	Port      string
	SecretKey string
}

func NewApp() (*App, error) {
	var err error
	app := App{
		Port:      os.Getenv("PORT"),
		SecretKey: os.Getenv("SECRET_KEY"),
	}
	err = app.validate()
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (a App) validate() error {
	return validation.ValidateStruct(
		&a,
		validation.Field(&a.Port, validation.Required),
		validation.Field(&a.SecretKey, validation.Required),
	)
}
