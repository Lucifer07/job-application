package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupServer(handlerUser *handler.UserHandler, handlerJob *handler.JobHandler, handlerapplied *handler.ApplicantHandler) *gin.Engine {
	router := gin.Default()
	router.ContextWithFallback = true
	router.Use(middleware.CustomMiddlewareError)
	router.POST("/login", handlerUser.Login)
	router.POST("/register", handlerUser.Register)
	adminRoute := router.Group("/admin")
	adminRoute.Use(middleware.MiddlewareJWTAuthorization, middleware.AuthorizationAdmin)
	routeAdminJob := adminRoute.Group("/jobs")
	routeAdminJob.POST("", handlerJob.CreateJob)
	routeAdminJob.PATCH("/quota", handlerJob.SetQuota)
	routeAdminJob.PATCH("/expired-date", handlerJob.SetExpiredDate)
	routeAdminJob.DELETE("/:id", handlerJob.CloseJob)
	userRoute := router.Group("/user")
	userRoute.Use(middleware.MiddlewareJWTAuthorization, middleware.AuthorizationUser)
	routeuserJob := userRoute.Group("/jobs")
	routeuserJob.POST("/:id", handlerapplied.AppliedJob)
	return router
}
