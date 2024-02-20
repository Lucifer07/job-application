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
	jobRepo:=repository.NewjobRepository(db)
	jobService:=service.NewJobService(jobRepo)
	userService := service.NewUserService(userRepo, &helperTool)
	jobHandler:=handler.NewJobHandler(jobService)
	UserHandler := handler.NewuserHandler(userService)
	router := server.SetupServer(UserHandler,jobHandler)
	if err := router.Run(":8080"); err != nil {
		log.Println(err)
	}
}
