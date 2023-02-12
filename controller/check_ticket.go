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
)

type (
	CheckTicketController struct {
		cu *usecase.CheckTicketUsecase
	}
)

func NewCheckTicketController(cu *usecase.CheckTicketUsecase) *CheckTicketController {
	return &CheckTicketController{cu: cu}
}

func (cc *CheckTicketController) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	ticketInput := model.CheckTicketRequest{}

	if err := c.Bind(&ticketInput); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}

	id := c.MustGet("id").(uint)

	if err := cc.cu.Create(ctx, &ticketInput, id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.ResponseWhenSuccess("success add ticket", nil))

}

func (cc *CheckTicketController) Mount(rg *gin.RouterGroup) {
	rg.POST("", middleware.ValidateJWToken(), cc.Create)
}
