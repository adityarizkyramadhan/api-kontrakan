package repository

import (
	"api-kontrakan/model"
	"context"

	"gorm.io/gorm"
)

type HouseRepository struct {
	db *gorm.DB
}

func NewHouseRepository(db *gorm.DB) *HouseRepository {
	return &HouseRepository{db: db}
}

func (hr *HouseRepository) Create(ctx context.Context, house *model.HouseModel) error {
	return hr.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(house).Error
	})
}

func (hr *HouseRepository) FindById(ctx context.Context, id uint) (*model.HouseModel, error) {
	house := new(model.HouseModel)
	err := hr.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Take(house).Error
	})
	return house, err
}
