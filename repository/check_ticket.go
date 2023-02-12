package repository

import (
	"api-kontrakan/model"
	"context"

	"gorm.io/gorm"
)

type (
	CheckTicketRepository struct {
		db *gorm.DB
	}
)

func NewCheckTicketRepository(db *gorm.DB) *CheckTicketRepository {
	return &CheckTicketRepository{db: db}
}

func (cr *CheckTicketRepository) Create(ctx context.Context, ticket *model.CheckTicketModel) error {
	return cr.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(ticket).Error
	})
}

func (cr *CheckTicketRepository) FindById(ctx context.Context, id uint) (*model.CheckTicketModel, error) {
	ticket := new(model.CheckTicketModel)
	err := cr.db.WithContext(ctx).Preload("User").Preload("House").Where("id = ?", id).Take(ticket).Error
	if err != nil {
		return nil, err
	}
	return ticket, err
}

func (cr *CheckTicketRepository) FindLastUserStatus(ctx context.Context, idHouse uint) ([]*model.CheckTicketModel, error) {
	var tickets []*model.CheckTicketModel
	err := cr.db.
		WithContext(ctx).
		Table("check_ticket_models").
		Select("check_ticket_models.*").
		Joins("LEFT JOIN (SELECT id_user, MAX(created_at) as created_at FROM check_ticket_models GROUP BY id_user) as t ON check_ticket_models.id_user = t.id_user AND check_ticket_models.created_at = t.created_at").
		Preload("User").
		Preload("House").
		Where("id_house = ?", idHouse).
		Scan(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, err
}
