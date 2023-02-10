package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	HouseModel struct {
		ID           uint `gorm:"primaryKey;autoIncrement"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
		DeletedAt    gorm.DeletedAt `gorm:"index"`
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

func (house *HouseModel) BeforeCreate(db *gorm.DB) error {
	house.CreatedAt = time.Now()
	house.UpdatedAt = time.Now()
	return nil
}

func (house *HouseModel) BeforeUpdate(db *gorm.DB) error {
	house.UpdatedAt = time.Now()
	return nil
}
