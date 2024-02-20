package main

import (
	"fmt"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/database"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/service"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	userRepo := repository.NewUserRepository(db)
	helperTool := util.HelperImpl{}
	userService := service.NewUserService(userRepo, &helperTool)
	UserHandler := handler.NewuserHandler(userService)
	router := server.SetupServer(UserHandler)
	if err := router.Run(":8080"); err != nil {
		log.Println(err)
	}
}
