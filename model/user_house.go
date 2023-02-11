package model

import "gorm.io/gorm"

type (
	UserHouseModel struct {
		gorm.Model
		IdUser  uint
		IdHouse uint
		User    UserModel  `gorm:"foreignkey:IdUser"`
		House   HouseModel `gorm:"foreignkey:IdHouse"`
	}
)
