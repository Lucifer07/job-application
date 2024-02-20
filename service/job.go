package service

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
)

type JobServiceImp struct {
	jobRepo repository.JobRepository
}
type JobService interface {
	CreateJob(ctx context.Context, job entity.Job) (*int, error)
	CloseJob(ctx context.Context, jobId int) error
	SetQuota(ctx context.Context, data dto.Quota) error
	SetExpiredDate(ctx context.Context, data dto.Expiry) error
}

func NewJobService(job repository.JobRepository) *JobServiceImp {
	return &JobServiceImp{
		jobRepo: job,
	}
}
func (s *JobServiceImp) CreateJob(ctx context.Context, job entity.Job) (*int, error) {
	id,err:=s.jobRepo.CreateJob(ctx,job)
	if err != nil {
		return nil, err
	}
	return id, nil
}
func (s *JobServiceImp) CloseJob(ctx context.Context, jobId int) error {
	jobs,err:=s.jobRepo.FindJob(ctx,jobId)
	if err != nil {
		return err
	}
	if jobs!=nil {
		err:=s.jobRepo.CloseJob(ctx,jobId)
		if err != nil {
			return err
		}
		return nil
	}
	return util.ErrorJobNotFound
}
func (s *JobServiceImp) SetQuota(ctx context.Context, data dto.Quota) error {
	jobs,err:=s.jobRepo.FindJob(ctx,data.Id)
	if err != nil {
		return err
	}
	if jobs!=nil {
		err:=s.jobRepo.SetQuota(ctx,data)
		if err != nil {
			return err
		}
		return nil
	}
	return util.ErrorJobNotFound
}
func (s *JobServiceImp) SetExpiredDate(ctx context.Context, data dto.Expiry) error {
	jobs,err:=s.jobRepo.FindJob(ctx,data.Id)
	if err != nil {
		return err
	}
	if jobs!=nil {
		err:=s.jobRepo.SetExpiredDate(ctx,data)
		if err != nil {
			return err
		}
		return nil
	}
	return util.ErrorJobNotFound
}

