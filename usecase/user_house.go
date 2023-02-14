package usecase

import "api-kontrakan/repository"

type (
	UserHouseUsecase struct {
		uhr *repository.UserHouseRepository
	}
)

func NewUserHouseUsecase(uhr *repository.UserHouseRepository) *UserHouseUsecase {
	return &UserHouseUsecase{uhr: uhr}
}
