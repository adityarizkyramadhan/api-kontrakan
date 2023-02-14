package usecase

import (
	"api-kontrakan/model"
	"api-kontrakan/repository"
	"context"
)

type (
	UserHouseUsecase struct {
		uhr *repository.UserHouseRepository
	}
)

func NewUserHouseUsecase(uhr *repository.UserHouseRepository) *UserHouseUsecase {
	return &UserHouseUsecase{uhr: uhr}
}

func (uhu *UserHouseUsecase) Create(ctx context.Context, input model.UserHouseRequest) error {
	userHouse := &model.UserHouseModel{
		IdUser:  input.IdUser,
		IdHouse: input.IdHouse,
	}
	return uhu.uhr.Create(ctx, userHouse)
}
