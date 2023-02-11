package model

import (
	"gorm.io/gorm"
)

func NewUserModel(ur UserRequestRegister) *UserModel {
	return &UserModel{
		Fullname: ur.Fullname,
		Username: ur.Username,
		Password: ur.Password,
	}
}

type (
	UserModel struct {
		gorm.Model
		Fullname     string
		Username     string `gorm:"unique"`
		Password     string
		CheckTickets []CheckTicketModel `gorm:"foreignkey:IdUser"`
		Houses       []HouseModel       `gorm:"many2many:user_houses;"`
	}

	UserRequestRegister struct {
		Fullname string `json:"Fullname"`
		Username string `json:"Username"`
		Password string `json:"Password"`
	}

	UserRequestLogin struct {
		Username string
		Password string
	}
)

func (UserModel) TableName() string {
	return "user_model"
}
