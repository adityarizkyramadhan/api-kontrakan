package config

import (
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Database struct {
	SupabaseUser     string
	SupabasePassword string
	SupabaseHost     string
	SupabasePort     string
	SupabaseDbName   string
}

func NewDatabase() (*Database, error) {
	var err error
	cfgDb := Database{
		SupabaseUser:     os.Getenv("SUPABASE_USER"),
		SupabasePassword: os.Getenv("SUPABASE_PASSWORD"),
		SupabaseHost:     os.Getenv("SUPABASE_HOST"),
		SupabasePort:     os.Getenv("SUPABASE_PORT"),
		SupabaseDbName:   os.Getenv("SUPABASE_DB_NAME"),
	}
	err = cfgDb.validate()
	if err != nil {
		return nil, err
	}
	return &cfgDb, err
}

func (db Database) validate() error {
	return validation.ValidateStruct(
		&db,
		validation.Field(&db.SupabaseUser, validation.Required),
		validation.Field(&db.SupabasePassword, validation.Required),
		validation.Field(&db.SupabaseHost, validation.Required),
		validation.Field(&db.SupabasePort, validation.Required),
		validation.Field(&db.SupabaseDbName, validation.Required),
	)
}

