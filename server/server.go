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
	routerAdminJob := adminRoute.Group("/jobs")
	routerAdminJob.GET("", handlerJob.FindAllJob)
	routerAdminJob.POST("", handlerJob.CreateJob)
	routerAdminJob.PATCH("/quota", handlerJob.SetQuota)
	routerAdminJob.PATCH("/expired-date", handlerJob.SetExpiredDate)
	routerAdminJob.DELETE("/:id", handlerJob.CloseJob)
	userRoute := router.Group("/user")
	userRoute.Use(middleware.MiddlewareJWTAuthorization, middleware.AuthorizationUser)
	routerUserJob := userRoute.Group("/jobs")
	routerUserJob.GET("", handlerJob.FindAllJob)
	routerUserJob.POST("/:id", handlerapplied.AppliedJob)
	return router
}
