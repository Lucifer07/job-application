package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/payload"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/service"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewuserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
func (h *UserHandler) Login(ctx *gin.Context) {
	var user dto.Login
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	jwt, err := h.userService.Login(ctx, payload.LoginToUser(user))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseMsg{Message: util.Success, Data: jwt})
}
func (h *UserHandler) Register(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	id,err := h.userService.Register(ctx, user)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseMsg{Message: util.Success,Data: response.Id{Id: *id}})
}
