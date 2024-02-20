package handler

import (
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/service"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
	"github.com/gin-gonic/gin"
)

type ApplicantHandler struct {
	applicantService service.AppicantRecordService
}

func NewApplicantHandler(ApplicantServices service.AppicantRecordService) *ApplicantHandler {
	return &ApplicantHandler{
		applicantService: ApplicantServices,
	}
}
func (h *ApplicantHandler) AppliedJob(ctx *gin.Context) {
	var appRecord dto.AppliedReq
	data := ctx.Value("data").(map[string]string)
	userId, err := strconv.Atoi(data["id"])
	if userId == 0 || err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	paramId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	appRecord.JobId = paramId
	appRecord.UserId = userId
	record, err := h.applicantService.AppliedJob(ctx, appRecord)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.ResponseMsg{Message: util.Success, Data: record})
}
