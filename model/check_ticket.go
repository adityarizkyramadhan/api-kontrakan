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
		IsCheckOut bool
		User       UserModel
		House      HouseModel
	}
)

func (CheckTicketModel) TableName() string {
	return "check_ticket"
}

func (checkTicket *CheckTicketModel) BeforeCreate(db *gorm.DB) error {
	checkTicket.CreatedAt = time.Now()
	return nil
}

func (checkTicket *CheckTicketModel) BeforeUpdate(db *gorm.DB) error {
	checkTicket.UpdatedAt = time.Now()
	return nil
}
