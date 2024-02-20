package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupServer(handlerUser *handler.UserHandler) *gin.Engine {
	router := gin.Default()
	router.ContextWithFallback = true
	router.Use(middleware.CustomMiddlewareError)
	router.POST("/login", handlerUser.Login)
	router.POST("/register", handlerUser.Register)
	return router
}
