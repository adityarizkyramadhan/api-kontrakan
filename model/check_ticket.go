package model

import (
	"gorm.io/gorm"
)

type (
	CheckTicketModel struct {
		gorm.Model
		IdUser     uint
		IdHouse    uint
		IsCheckOut bool `gorm:"default:false"`
		Deskripsi  string
		User       UserModel  `gorm:"foreignkey:IdUser;association_foreignkey:ID"`
		House      HouseModel `gorm:"foreignkey:IdHouse;association_foreignkey:ID"`
	}
)

func (CheckTicketModel) TableName() string {
	return "check_ticket"
}
