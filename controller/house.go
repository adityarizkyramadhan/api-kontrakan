package controller

import (
	"api-kontrakan/model"
	"api-kontrakan/usecase"
	"api-kontrakan/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	HouseController struct {
		hu *usecase.HouseUsecase
	}
)

func NewHouseController(hu *usecase.HouseUsecase) *HouseController {
	return &HouseController{hu: hu}
}

func (hc *HouseController) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	houseInput := model.HouseRequestCreate{}

	if err := c.Bind(&houseInput); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
}
