package usecase

import (
	"api-kontrakan/model"
	"api-kontrakan/repository"
	"context"
)

type (
	CheckTicketUsecase struct {
		cr *repository.CheckTicketRepository
	}
)

func NewCheckTicketUsecase(cr *repository.CheckTicketRepository) *CheckTicketUsecase {
	return &CheckTicketUsecase{cr: cr}
}

func (cu *CheckTicketUsecase) Create(ctx context.Context, input *model.CheckTicketRequest, idUser uint) error {
	ticket := &model.CheckTicketModel{
		IdUser:    idUser,
		IdHouse:   input.IdHouse,
		Deskripsi: input.Deskripsi,
	}
	return cu.cr.Create(ctx, ticket)
}
