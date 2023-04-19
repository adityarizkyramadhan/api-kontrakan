package controller

import (
	"api-kontrakan/middleware"
	"api-kontrakan/model"
	"api-kontrakan/usecase"
	"api-kontrakan/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

	if err := hc.hu.CreateHouse(ctx, &houseInput); err != nil {
		if errors.Cause(err) == utils.ErrValidation {
			c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error()))
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.ResponseWhenSuccess("success create house", nil))
}

func (hc *HouseController) FindById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	house, err := hc.hu.FindById(ctx, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success find by id house", house))
}

func (hc *HouseController) Mount(rg *gin.RouterGroup) {
	rg.GET("details/:id", middleware.ValidateJWToken(), hc.FindById)
	rg.POST("", middleware.ValidateJWToken(), hc.Create)
}
