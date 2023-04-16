package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
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
		ID           uint      `gorm:"primary_key;auto_increment"`
		CreatedAt    time.Time `gorm:"autoCreateTime"`
		UpdatedAt    time.Time `gorm:"autoUpdateTime"`
		Fullname     string
		Username     string `gorm:"unique"`
		Password     string
		CheckTickets []CheckTicketModel `gorm:"foreignkey:IdUser"`
		Houses       []HouseModel       `gorm:"many2many:user_houses;"`
	}

	UserRequestRegister struct {
		Fullname string `json:"Fullname"`
		Username string `json:"Username"`
		Password string `json:"-"`
	}

	UserRequestLogin struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
	}
)

func (UserModel) TableName() string {
	return "user_model"
}

func (r UserRequestRegister) Validate() error {
	return validation.ValidateStruct(&r,
		// Fullname is required and should not be empty
		validation.Field(&r.Fullname, validation.Required, validation.Length(1, 100)),
		// Username is required and should be a valid email address
		validation.Field(&r.Username, validation.Required, is.Email),
		// Password is required and should be at least 6 characters long
		validation.Field(&r.Password, validation.Required, validation.Length(6, 100)),
	)
}

func (r UserRequestLogin) Validate() error {
	return validation.ValidateStruct(&r,
		// Username is required and should not be empty
		validation.Field(&r.Username, validation.Required),
		// Password is required and should not be empty
		validation.Field(&r.Password, validation.Required),
	)
}
