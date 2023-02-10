package main

import (
	"api-kontrakan/config"
	"api-kontrakan/controller"
	"api-kontrakan/repository"
	"api-kontrakan/usecase"
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
	//user
	repoUser := repository.NewUserRepository(db)
	usecaseUser := usecase.NewUserusecase(repoUser)
	ctrlUser := controller.NewUserController(usecaseUser)
	groupUser := r.Group("user")
	ctrlUser.Mount(groupUser)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
