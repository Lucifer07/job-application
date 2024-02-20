package service

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
)

type JobServiceImp struct {
	jobRepo     repository.JobRepository
	transaction util.Transactor
}
type JobService interface {
	CreateJob(ctx context.Context, job entity.Job) (*int, error)
	CloseJob(ctx context.Context, jobId int) error
	SetQuota(ctx context.Context, data dto.Quota) error
	SetExpiredDate(ctx context.Context, data dto.Expiry) error
}

func NewJobService(job repository.JobRepository, transaction util.Transactor) *JobServiceImp {
	return &JobServiceImp{
		jobRepo:     job,
		transaction: transaction,
	}
}
func (s *JobServiceImp) CreateJob(ctx context.Context, job entity.Job) (*int, error) {
	id, err := s.jobRepo.CreateJob(ctx, job)
	if err != nil {
		return nil, err
	}
	return id, nil
}
func (s *JobServiceImp) CloseJob(ctx context.Context, jobId int) error {
	err := s.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		jobs, err := s.jobRepo.FindJob(ctx, jobId)
		if err != nil {
			return err
		}
		if jobs != nil {
			err := s.jobRepo.CloseJob(ctx, jobId)
			if err != nil {
				return err
			}
			return nil
		}
		return util.ErrorJobNotFound
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *JobServiceImp) SetQuota(ctx context.Context, data dto.Quota) error {
	err := s.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		jobs, err := s.jobRepo.FindJob(ctx, data.Id)
		if err != nil {
			return err
		}
		if jobs != nil {
			err := s.jobRepo.SetQuota(ctx, data)
			if err != nil {
				return err
			}
			return nil
		}
		return util.ErrorJobNotFound
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *JobServiceImp) SetExpiredDate(ctx context.Context, data dto.Expiry) error {
	err := s.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		jobs, err := s.jobRepo.FindJob(ctx, data.Id)
		if err != nil {
			return err
		}
		if jobs != nil {
			err := s.jobRepo.SetExpiredDate(ctx, data)
			if err != nil {
				return err
			}
			return nil
		}
		return util.ErrorJobNotFound
	})
	if err != nil {
		return err
	}
	return nil
}
