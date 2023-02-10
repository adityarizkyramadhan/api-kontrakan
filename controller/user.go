package controller

import (
	"api-kontrakan/model"
	"api-kontrakan/usecase"
	"api-kontrakan/utils"
	"net/http"

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
	ctx := c.Request.Context()
	userInput := new(model.UserRequestRegister)
	if err := c.BindJSON(userInput); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	data, err := uc.uu.SearchByUsername(ctx, userInput.Username)
	if err == nil && data != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(utils.ErrUniqueUsername.Error()))
		return
	}
	token, err := uc.uu.Register(ctx, userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseWhenSuccess("success register user", gin.H{"Token": token}))
}

func (uc *UserController) Login(c *gin.Context) {
	ctx := c.Request.Context()
	userInput := new(model.UserRequestLogin)
	if err := c.BindJSON(userInput); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	_, err := uc.uu.SearchByUsername(ctx, userInput.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error()))
		return
	}
	token, err := uc.uu.Login(ctx, userInput)
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
