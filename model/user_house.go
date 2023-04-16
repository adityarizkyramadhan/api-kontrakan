package model

import "time"

type (
	UserHouseModel struct {
		ID        uint      `gorm:"primary_key;auto_increment"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
		IdUser    uint
		IdHouse   uint
		User      UserModel  `gorm:"foreignkey:IdUser"`
		House     HouseModel `gorm:"foreignkey:IdHouse"`
	}

	UserHouseRequest struct {
		IdUser  uint
		IdHouse uint
	}
)
