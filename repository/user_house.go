package repository

import (
	"api-kontrakan/model"
	"context"

	"gorm.io/gorm"
)

type (
	UserHouseRepository struct {
		db *gorm.DB
	}
)

func NewUserHouseRepository(db *gorm.DB) *UserHouseRepository {
	return &UserHouseRepository{db: db}
}

func (uhr *UserHouseRepository) Create(ctx context.Context, userHouse *model.UserHouseModel) error {
	return uhr.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(userHouse).Error
	})
}

func (uhr *UserHouseRepository) FindHouseByUser(ctx context.Context, idUser uint) ([]*model.UserHouseModel, error) {
	var userHouses []*model.UserHouseModel
	err := uhr.db.Preload("House").Preload("User").Where("id_user = ?", idUser).Find(&userHouses).Error
	if err != nil {
		return nil, err
	}
	return userHouses, nil
}

func (uhr *UserHouseRepository) FindUserByHouse(ctx context.Context, idHouse uint) ([]*model.UserHouseModel, error) {
	var userHouses []*model.UserHouseModel
	err := uhr.db.Preload("House").Preload("User").Where("id_house = ?", idHouse).Find(&userHouses).Error
	if err != nil {
		return nil, err
	}
	return userHouses, nil
}
