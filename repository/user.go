package repository

import (
	"api-kontrakan/model"
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *model.UserModel) error {
	return ur.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(user).Error
	})
}

func (ur *UserRepository) SearchByUsername(ctx context.Context, username string) (*model.UserModel, error) {
	user := new(model.UserModel)
	err := ur.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Where("username = ?", username).Take(user).Error
	})
	return user, err
}

func (ur *UserRepository) SearchByID(ctx context.Context, id uint) (*model.UserModel, error) {
	user := new(model.UserModel)
	err := ur.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Take(user).Error
	})
	return user, err
}
