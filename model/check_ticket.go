package model

import "time"

type (
	CheckTicketModel struct {
		ID         uint      `gorm:"primary_key;auto_increment"`
		CreatedAt  time.Time `gorm:"autoCreateTime"`
		UpdatedAt  time.Time `gorm:"autoUpdateTime"`
		IdUser     uint
		IdHouse    uint
		IsCheckOut bool `gorm:"default:false"`
		Deskripsi  string
		User       UserModel  `gorm:"foreignkey:IdUser;association_foreignkey:ID"`
		House      HouseModel `gorm:"foreignkey:IdHouse;association_foreignkey:ID"`
	}

	CheckTicketRequest struct {
		IdHouse   uint
		Deskripsi string
	}
)

func (CheckTicketModel) TableName() string {
	return "check_ticket"
}
