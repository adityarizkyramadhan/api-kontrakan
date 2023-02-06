package config

import (
	"api-kontrakan/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MakeConnectionDatabase(data Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s "+
		"password=%s "+
		"host=%s "+
		"TimeZone=Asia/Singapore "+
		"port=%s "+
		"dbname=%s", data.SupabaseUser, data.SupabasePassword, data.SupabaseHost, data.SupabasePort, data.SupabaseDbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(new(model.UserModel), new(model.HouseModel), new(model.CheckTicketModel)); err != nil {
		return nil, err
	}
	return db, nil
}
