package model

import "time"

type (
	HouseModel struct {
		ID           uint   `gorm:"primary_key;auto_increment"`
		CreatedAt    time.Time `gorm:"autoCreateTime"`
		UpdatedAt    time.Time `gorm:"autoUpdateTime"`
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
