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
	UserController struct {
		uu *usecase.UserUsecase
	}
)

func NewUserController(uu *usecase.UserUsecase) *UserController {
	return &UserController{uu: uu}
}

func (uc *UserController) Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	userInput := model.UserRequestRegister{}
	if err := c.Bind(&userInput); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	token, err := uc.uu.Register(ctx, &userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseWhenSuccess("success register user", gin.H{"Token": token}))
}

func (uc *UserController) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	userInput := model.UserRequestLogin{}
	if err := c.Bind(&userInput); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	token, err := uc.uu.Login(ctx, &userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseWhenSuccess("success login user", gin.H{"Token": token}))
}

func (uc *UserController) Mount(rg *gin.RouterGroup) {
	rg.POST("login", uc.Login)
	rg.POST("register", uc.Register)
}
