package usecase_test

import (
	"api-kontrakan/mocks"
	"api-kontrakan/model"
	"api-kontrakan/usecase"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateHouse(t *testing.T) {
	tests := []struct {
		name          string
		inputHouse    *model.HouseRequestCreate
		dbHouse       *model.HouseModel
		expectedError error
	}{
		// TODO: Add test cases.
		{
			name: "Success add house by admin",
			inputHouse: &model.HouseRequestCreate{
				NamaAlamat:  "Jalan Semangka No 54, Sukamiskin",
				NamaSebutan: "Rumah Koruptor",
				Deskripsi:   "Ini Deskripsi",
			},
			dbHouse: &model.HouseModel{
				NamaAlamat:   "Jalan Semangka No 54, Sukamiskin",
				NamaSebutan:  "Rumah Koruptor",
				Deskripsi:    "Ini Deskripsi",
				CheckTickets: nil,
				Users:        nil,
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoHouse := mocks.NewHouseRepositoryImplementation(t)
			repoHouse.Mock.On("Create", context.TODO(), tt.dbHouse).Return(nil)
			usecaseHouse := usecase.NewHouseUsecase(repoHouse)
			err := usecaseHouse.CreateHouse(context.TODO(), tt.inputHouse)
			if tt.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError.Error())
			}
		})
	}
}
