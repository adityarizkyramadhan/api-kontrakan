package usecase

import (
	"api-kontrakan/middleware"
	"api-kontrakan/model"
	"api-kontrakan/repository"
	"api-kontrakan/utils"
	"context"
)

type UserUsecase struct {
	ur repository.UserRepositoryImplementation
}

func NewUserusecase(ur repository.UserRepositoryImplementation) *UserUsecase {
	return &UserUsecase{ur: ur}
}

func (uc *UserUsecase) Register(ctx context.Context, input *model.UserRequestRegister) (string, error) {
	user := &model.UserModel{
		Fullname: input.Fullname,
		Password: input.Password,
		Username: input.Username,
	}
	err := uc.ur.Create(ctx, user)
	if err != nil {
		return "", err
	}
	token, err := middleware.GenerateJWToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uc *UserUsecase) SearchByID(ctx context.Context, id int) (*model.UserModel, error) {
	if id <= 0 {
		return nil, utils.ErrId
	}
	return uc.ur.SearchByID(ctx, uint(id))
}

func (uc *UserUsecase) SearchByUsername(ctx context.Context, username string) (*model.UserModel, error) {
	if username == "" {
		return nil, utils.ErrUsername
	}
	return uc.SearchByUsername(ctx, username)
}
