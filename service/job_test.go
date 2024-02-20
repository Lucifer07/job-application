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

func TestJobServiceImp_CreateJob(t *testing.T) {
	t.Run("return 200 if succes create job", func(t *testing.T) {
		// given
		ctx:=new(context.Context)
		idtarget := 1
		jobRepo := new(mocks.JobRepository)
		transactor:=new(mocks.Transactor)
		job := entity.Job{}
		jobService:=service.NewJobService(jobRepo,transactor)
		// when
		jobRepo.On("CreateJob", mock.Anything, mock.Anything).Return(&idtarget, nil)
		id,_:=jobService.CreateJob(*ctx,job)
		assert.Equal(t,idtarget,*id)
	})
	t.Run("return error if failed create job", func(t *testing.T) {
		// given
		ctx:=new(context.Context)
		jobRepo := new(mocks.JobRepository)
		transactor:=new(mocks.Transactor)
		job := entity.Job{}
		jobService:=service.NewJobService(jobRepo,transactor)
		// when
		jobRepo.On("CreateJob", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		_,err:=jobService.CreateJob(*ctx,job)
		assert.Equal(t,util.ErrorInternal,err)
	})
}
