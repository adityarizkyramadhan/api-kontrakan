package model

import (
	"gorm.io/gorm"
)

type (
	HouseModel struct {
		gorm.Model
		NamaAlamat   string
		NamaSebutan  string
		Deskripsi    string
		CheckTickets []CheckTicketModel `gorm:"foreignkey:IdHouse"`
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
