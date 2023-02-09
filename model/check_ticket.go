package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	CheckTicketModel struct {
		ID         uint `gorm:"primaryKey;autoIncrement"`
		CreatedAt  time.Time
		UpdatedAt  time.Time
		DeletedAt  gorm.DeletedAt `gorm:"index"`
		IdUser     uint
		IdRumah    uint
		IsCheckOut bool `gorm:"default:'false'"`
		Deskripsi  string
		User       UserModel  `gorm:"foreignkey:IdUser;association_foreignkey:ID"`
		House      HouseModel `gorm:"foreignkey:IdHouse;association_foreignkey:ID"`
	}
)

func (CheckTicketModel) TableName() string {
	return "check_ticket"
}

func (checkTicket *CheckTicketModel) BeforeCreate(db *gorm.DB) error {
	checkTicket.CreatedAt = time.Now()
	checkTicket.UpdatedAt = time.Now()
	return nil
}

func (checkTicket *CheckTicketModel) BeforeUpdate(db *gorm.DB) error {
	checkTicket.UpdatedAt = time.Now()
	return nil
}
