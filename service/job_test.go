package service_test

import (
	"context"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/service"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewJobService(t *testing.T) {
	t.Run("should return Job service", func(t *testing.T) {
		// given
		targetService := service.JobServiceImp{}
		jobRepo := new(mocks.JobRepository)
		transactor := new(mocks.Transactor)
		// when
		jobService := service.NewJobService(jobRepo, transactor)
		assert.IsType(t, targetService, *jobService)
	})
}
func TestJobServiceImp_CreateJob(t *testing.T) {
	t.Run("return id if succes create job", func(t *testing.T) {
		// given
		ctx := new(context.Context)
		idtarget := 1
		jobRepo := new(mocks.JobRepository)
		transactor := new(mocks.Transactor)
		job := entity.Job{}
		jobService := service.NewJobService(jobRepo, transactor)
		// when
		jobRepo.On("CreateJob", mock.Anything, mock.Anything).Return(&idtarget, nil)
		id, _ := jobService.CreateJob(*ctx, job)
		// then
		assert.Equal(t, idtarget, *id)
	})
	t.Run("return error if failed create job", func(t *testing.T) {
		// given
		ctx := new(context.Context)
		jobRepo := new(mocks.JobRepository)
		transactor := new(mocks.Transactor)
		job := entity.Job{}
		jobService := service.NewJobService(jobRepo, transactor)
		// when
		jobRepo.On("CreateJob", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		_, err := jobService.CreateJob(*ctx, job)
		// then
		assert.Equal(t, util.ErrorInternal, err)
	})
}

func TestJobServiceImp_FindAllJob(t *testing.T) {
	t.Run("should return slice of job if success", func(t *testing.T) {
		// given
		ctx := new(context.Context)
		dataTarget := []entity.Job{}
		jobRepo := new(mocks.JobRepository)
		transactor := new(mocks.Transactor)
		jobService := service.NewJobService(jobRepo, transactor)
		// when
		jobRepo.On("FindAllJob", mock.Anything, mock.Anything).Return(dataTarget, nil)
		data, _ := jobService.FindAllJob(*ctx)
		// then
		assert.Equal(t, dataTarget, data)
	})
	t.Run("should err  if have internal server error", func(t *testing.T) {
		// given
		ctx := new(context.Context)
		jobRepo := new(mocks.JobRepository)
		transactor := new(mocks.Transactor)
		jobService := service.NewJobService(jobRepo, transactor)
		// when
		jobRepo.On("FindAllJob", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		_, err := jobService.FindAllJob(*ctx)
		// then
		assert.Equal(t, util.ErrorInternal, err)
	})
}

func TestJobServiceImp_CloseJob(t *testing.T) {
	t.Run("should return nil if successed", func(t *testing.T) {
		// given
		job := entity.Job{}
		jobId := 1
		ctx := new(context.Context)
		jobRepo := new(mocks.JobRepository)
		transactor := new(mocks.Transactor)
		jobService := service.NewJobService(jobRepo, transactor)
		// when
		jobRepo.On("FindJob", mock.Anything, mock.Anything).Return(job, nil)
		jobRepo.On("CloseJob", mock.Anything, mock.Anything).Return(nil)
		transactor.On("WithinTransaction", mock.Anything, mock.Anything).Return(nil)
		err := jobService.CloseJob(*ctx, jobId)
		// then
		assert.Equal(t, nil, err)
	})
	t.Run("should return error job not found if cant find job id", func(t *testing.T) {
		// given
		jobId := 1
		ctx := new(context.Context)
		jobRepo := new(mocks.JobRepository)
		transactor := new(mocks.Transactor)
		jobService := service.NewJobService(jobRepo, transactor)
		// when
		jobRepo.On("FindJob", mock.Anything, mock.Anything).Return(nil, nil)
		transactor.On("WithinTransaction", mock.Anything, mock.Anything).Return(util.ErrorJobNotFound)
		err := jobService.CloseJob(*ctx, jobId)
		// then
		assert.Equal(t, util.ErrorJobNotFound, err)
	})

	t.Run("should return error internal server if have problem in query when find job", func(t *testing.T) {
		// given
		jobId := 1
		ctx := new(context.Context)
		jobRepo := new(mocks.JobRepository)
		transactor := new(mocks.Transactor)
		jobService := service.NewJobService(jobRepo, transactor)
		// when
		jobRepo.On("FindJob", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		transactor.On("WithinTransaction", mock.Anything, mock.Anything).Return(util.ErrorInternal)
		err := jobService.CloseJob(*ctx, jobId)
		// then
		assert.Equal(t, util.ErrorInternal, err)
	})
	t.Run("should return error internal server if have problem in query when try to close job", func(t *testing.T) {
		// given
		jobId := 1
		ctx := new(context.Context)
		jobRepo := new(mocks.JobRepository)
		transactor := new(mocks.Transactor)
		jobService := service.NewJobService(jobRepo, transactor)
		// when
		jobRepo.On("FindJob", mock.Anything, mock.Anything).Return(nil, nil)
		jobRepo.On("CloseJob", mock.Anything, mock.Anything).Return(util.ErrorInternal)
		transactor.On("WithinTransaction", mock.Anything, mock.Anything).Return(util.ErrorInternal)
		err := jobService.CloseJob(*ctx, jobId)
		// then
		assert.Equal(t, util.ErrorInternal, err)
	})
}
