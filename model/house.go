package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type (
	HouseModel struct {
		gorm.Model
		NamaAlamat   string
		NamaSebutan  string
		Deskripsi    string
		CheckTickets []CheckTicketModel `gorm:"foreignkey:IdHouse"`
		Users        []UserModel        `gorm:"many2many:user_houses;"`
	}

	HouseRequestCreate struct {
		NamaAlamat  string
		NamaSebutan string
		Deskripsi   string
	}
)

func (HouseModel) TableName() string {
	return "house_model"
}

func (r HouseRequestCreate) Validate() error {
	return validation.ValidateStruct(&r,
		// NamaAlamat is required and should not be empty
		validation.Field(&r.NamaAlamat, validation.Required),
		// NamaSebutan is required and should not be empty
		validation.Field(&r.NamaSebutan, validation.Required),
		// Deskripsi is optional but should not be longer than 1000 characters
		validation.Field(&r.Deskripsi, validation.Length(0, 1000)),
	)
}
