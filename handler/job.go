package handler

import (
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/service"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	JobService service.JobService
}

func NewJobHandler(JobService service.JobService) *JobHandler {
	return &JobHandler{
		JobService: JobService,
	}
}
func (h *JobHandler) CreateJob(ctx *gin.Context) {
	var job entity.Job
	if err := ctx.ShouldBindJSON(&job); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	id, err := h.JobService.CreateJob(ctx, job)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.ResponseMsg{Message: util.Success, Data: response.Id{Id: *id}})
}
func (h *JobHandler) CloseJob(ctx *gin.Context) {
	paramId := ctx.Param("id")
	jobId, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	err = h.JobService.CloseJob(ctx, jobId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusNoContent, util.NoContent)
}
func (h *JobHandler) SetQuota(ctx *gin.Context) {
	var job dto.Quota
	if err := ctx.ShouldBindJSON(&job); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	err := h.JobService.SetQuota(ctx, job)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusNoContent, util.NoContent)
}
func (h *JobHandler) SetExpiredDate(ctx *gin.Context) {
	var job dto.Expiry
	if err := ctx.ShouldBindJSON(&job); err != nil {
		ctx.Error(util.ErrorBadRequest)
		return
	}
	err := h.JobService.SetExpiredDate(ctx, job)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusNoContent, util.NoContent)
}
