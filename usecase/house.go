package usecase

import (
	"api-kontrakan/model"
	"api-kontrakan/repository"
	"context"
)

type (
	HouseUsecase struct {
		hr repository.HouseRepositoryImplementation
	}
)

func NewHouseUsecase(hr repository.HouseRepositoryImplementation) *HouseUsecase {
	return &HouseUsecase{hr: hr}
}

func (hu *HouseUsecase) CreateHouse(ctx context.Context, input model.HouseRequestCreate) error {
	house := &model.HouseModel{
		NamaAlamat:  input.NamaAlamat,
		NamaSebutan: input.NamaSebutan,
		Deskripsi:   input.Deskripsi,
	}
	return hu.hr.Create(ctx, house)
}
