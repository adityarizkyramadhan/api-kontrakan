package usecase

import (
	"api-kontrakan/middleware"
	"api-kontrakan/model"
	"api-kontrakan/repository"
	"context"
)

type UserUsecase struct {
	ur repository.UserRepositoryImplementation
}

func NewUserusecase(ur repository.UserRepositoryImplementation) *UserUsecase {
	return &UserUsecase{ur: ur}
}

func (uc *UserUsecase) Register(ctx context.Context, input model.UserRequestRegister) (string, error) {
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
