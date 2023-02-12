package usecase

import (
	"api-kontrakan/model"
	"api-kontrakan/repository"
	"context"
	"strconv"
)

type (
	HouseUsecase struct {
		hr repository.HouseRepositoryImplementation
	}
)

func NewHouseUsecase(hr repository.HouseRepositoryImplementation) *HouseUsecase {
	return &HouseUsecase{hr: hr}
}

func (hu *HouseUsecase) CreateHouse(ctx context.Context, input *model.HouseRequestCreate) error {
	house := &model.HouseModel{
		NamaAlamat:  input.NamaAlamat,
		NamaSebutan: input.NamaSebutan,
		Deskripsi:   input.Deskripsi,
	}
	return hu.hr.Create(ctx, house)
}

func (hu *HouseUsecase) FindById(ctx context.Context, id string) (*model.HouseModel, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return hu.hr.FindById(ctx, uint(idInt))
}
