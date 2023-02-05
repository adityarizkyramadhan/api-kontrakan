package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	HouseModel struct {
		ID          uint `gorm:"primaryKey;autoIncrement"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   gorm.DeletedAt `gorm:"index"`
		NamaAlamat  string
		NamaSebutan string
	}
)
