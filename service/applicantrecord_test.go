package service_test

import (
	"context"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/service"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAppicantRecordService(t *testing.T) {
	// given
	targetService := service.AppicantRecordImp{}
	transactor := new(mocks.Transactor)
	appRepo := new(mocks.ApplicantRecordRepository)
	jobRepo := new(mocks.JobRepository)
	// when
	service := service.NewAppicantRecordService(appRepo, jobRepo, transactor)
	// then
	assert.IsType(t, targetService, *service)
}

func TestAppicantRecordImp_AppliedJob(t *testing.T) {
	t.Run("should return error nil if succes applied", func(t *testing.T) {
		// given
		appRec:=dto.AppliedReq{}
		ctx:=context.Background()
		transactor := new(mocks.Transactor)
		appRepo := new(mocks.ApplicantRecordRepository)
		jobRepo := new(mocks.JobRepository)
		service := service.NewAppicantRecordService(appRepo, jobRepo, transactor)
		// when
		transactor.On("WithinTransaction",mock.Anything,mock.Anything).Return(nil)
		_,err:=service.AppliedJob(ctx,appRec)
		assert.Equal(t,nil,err)

	})
	t.Run("should return error internal server error if have problem in transaction", func(t *testing.T) {
		// given
		appRec:=dto.AppliedReq{}
		ctx:=context.Background()
		transactor := new(mocks.Transactor)
		appRepo := new(mocks.ApplicantRecordRepository)
		jobRepo := new(mocks.JobRepository)
		service := service.NewAppicantRecordService(appRepo, jobRepo, transactor)
		// when
		transactor.On("WithinTransaction",mock.Anything,mock.Anything).Return(util.ErrorInternal)
		_,err:=service.AppliedJob(ctx,appRec)
		assert.Equal(t,util.ErrorInternal,err)

	})
}
