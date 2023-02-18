package main

import (
	"api-kontrakan/config"
	"api-kontrakan/controller"
	"api-kontrakan/middleware"
	"api-kontrakan/repository"
	"api-kontrakan/usecase"
	"api-kontrakan/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	databaseConf, err := config.NewDatabase()
	if err != nil {
		panic(err.Error())
	}
	db, err := config.MakeConnectionDatabase(databaseConf)
	if err != nil {
		panic(err.Error())
	}
	r := gin.Default()
	r.Use(middleware.TimeoutMiddleware())
	r.GET("health", func(c *gin.Context) {
		c.JSON(200, utils.ResponseWhenSuccess("success", "deploy health 100% after restar"))
	})

	//user
	repoUser := repository.NewUserRepository(db)
	usecaseUser := usecase.NewUserusecase(repoUser)
	ctrlUser := controller.NewUserController(usecaseUser)
	groupUser := r.Group("user")
	ctrlUser.Mount(groupUser)

	//house
	repoHouse := repository.NewHouseRepository(db)
	usecaseHouse := usecase.NewHouseUsecase(repoHouse)
	ctrlHouse := controller.NewHouseController(usecaseHouse)
	groupHouse := r.Group("house")
	ctrlHouse.Mount(groupHouse)

	//check ticket
	repoTicket := repository.NewCheckTicketRepository(db)
	usecaseTicket := usecase.NewCheckTicketUsecase(repoTicket)
	ctrlTicket := controller.NewCheckTicketController(usecaseTicket)
	groupTicket := r.Group("ticket")
	ctrlTicket.Mount(groupTicket)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
