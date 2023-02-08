package model

import (
	"time"

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
		ID           uint `gorm:"primaryKey;autoIncrement"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
		DeletedAt    gorm.DeletedAt `gorm:"index"`
		Fullname     string
		Username     string `gorm:"unique"`
		Password     string
		CheckTickets []CheckTicketModel `gorm:"foreignkey:IdUser"`
	}

	UserRequestRegister struct {
		Fullname string
		Username string
		Password string
	}

	UserRequestLogin struct {
		Username string
		Password string
	}
)

func (UserModel) TableName() string {
	return "user_model"
}

func (user *UserModel) BeforeCreate(db *gorm.DB) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return nil
}

func (user *UserModel) BeforeUpdate(db *gorm.DB) error {
	user.UpdatedAt = time.Now()
	return nil
}
