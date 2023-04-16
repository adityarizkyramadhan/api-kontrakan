package usecase

import (
	"api-kontrakan/middleware"
	"api-kontrakan/model"
	"api-kontrakan/repository"
	"api-kontrakan/utils"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	ur repository.UserRepositoryImplementation
}

func NewUserusecase(ur repository.UserRepositoryImplementation) *UserUsecase {
	return &UserUsecase{ur: ur}
}

func (uc *UserUsecase) Register(ctx context.Context, input *model.UserRequestRegister) (string, error) {
	temp, err := uc.SearchByUsername(ctx, input.Username)
	if err == nil && temp != nil {
		return "", utils.ErrUniqueUsername
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		return "", err
	}
	user := &model.UserModel{
		Fullname: input.Fullname,
		Password: string(hashPassword),
		Username: input.Username,
	}
	err = uc.ur.Create(ctx, user)
	if err != nil {
		return "", err
	}
	token, err := middleware.GenerateJWToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uc *UserUsecase) SearchByID(ctx context.Context, id uint) (*model.UserModel, error) {
	if id <= 0 {
		return nil, utils.ErrId
	}
	return uc.ur.SearchByID(ctx, id)
}

func (uc *UserUsecase) SearchByUsername(ctx context.Context, username string) (*model.UserModel, error) {
	if username == "" {
		return nil, utils.ErrUsername
	}
	return uc.ur.SearchByUsername(ctx, username)
}

func (uc *UserUsecase) Login(ctx context.Context, input *model.UserRequestLogin) (string, error) {
	user, err := uc.SearchByUsername(ctx, input.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", err
	}
	token, err := middleware.GenerateJWToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
