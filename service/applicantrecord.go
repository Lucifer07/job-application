package service

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
)

type AppicantRecordImp struct {
	AppicantRecordRepo repository.ApplicantRecordRepository
	jobRepo            repository.JobRepository
	transaction        util.Transactor
}
type AppicantRecordService interface {
	AppliedJob(ctx context.Context, AppicantRecord dto.AppliedReq) (*dto.ApplicantRecord, error)
}

func NewAppicantRecordService(AppicantRecord repository.ApplicantRecordRepository, jobRepo repository.JobRepository, transaction util.Transactor) *AppicantRecordImp {
	return &AppicantRecordImp{
		AppicantRecordRepo: AppicantRecord,
		jobRepo:            jobRepo,
		transaction:        transaction,
	}
}
func (u *AppicantRecordImp) AppliedJob(ctx context.Context, AppicantRecord dto.AppliedReq) (*dto.ApplicantRecord, error) {
	var result *dto.ApplicantRecord
	err := u.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		var dataJob dto.Quota
		dataJob.Id = AppicantRecord.JobId
		job, err := u.jobRepo.FindJob(ctx, AppicantRecord.JobId)
		if err != nil {
			return err
		}
		if job != nil {
			if job.Quota <= 0 {
				return util.ErrorJobfull
			}
			dataJob.Quota = job.Quota - 1
			record, err := u.AppicantRecordRepo.FindAppliedJob(ctx, AppicantRecord.JobId, AppicantRecord.UserId)
			if err != nil {
				return err
			}
			if record != nil {
				return util.ErrorAlreadyApplied
			}
			_, err = u.AppicantRecordRepo.AppliedJob(ctx, AppicantRecord)
			if err != nil {
				return err
			}
			err = u.jobRepo.SetQuota(ctx, dataJob)
			if err != nil {
				return err
			}
			record, _ = u.AppicantRecordRepo.FindAppliedJob(ctx, AppicantRecord.JobId, AppicantRecord.UserId)
			result = record
			return nil
		}
		return util.ErrorJobNotFound
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
