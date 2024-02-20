package repository

import (
	"context"
	"database/sql"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
)

type JobRepositoryDb struct {
	db *sql.DB
}
type JobRepository interface {
	CreateJob(ctx context.Context, job entity.Job) (*int, error)
	CloseJob(ctx context.Context, jobId int) error
	SetQuota(ctx context.Context, data dto.Quota) error
	SetExpiredDate(ctx context.Context, data dto.Expiry) error
	FindJob(ctx context.Context, jobId int) (*entity.Job, error)
}

func NewjobRepository(db *sql.DB) *JobRepositoryDb {
	return &JobRepositoryDb{
		db: db,
	}
}
func (r *JobRepositoryDb) FindJob(ctx context.Context, jobId int) (*entity.Job, error) {
	var data entity.Job
	statment := `select id,name,quota,expired_date from jobs where deleted_at is null and id =$1;`
	err := r.db.QueryRowContext(ctx, statment, jobId).Scan(&data.Id, &data.Name, &data.Quota, &data.ExpiredDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}
func (r *JobRepositoryDb) CreateJob(ctx context.Context, job entity.Job) (*int, error) {
	var id int
	statment := `insert into jobs(name,quota,expired_date)values($1,$1,$3) returning id;`
	err := r.db.QueryRowContext(ctx, statment, job.Name, job.Quota, job.ExpiredDate).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
func (r *JobRepositoryDb) CloseJob(ctx context.Context, jobId int) error {
	statment := `update jobs set deleted_at=now() where id=$1 and deleted_at is not null`
	_, err := r.db.ExecContext(ctx, statment, jobId)
	if err != nil {
		return err
	}
	return nil
}
func (r *JobRepositoryDb) SetQuota(ctx context.Context, data dto.Quota) error {
	statment := `update jobs set quota=$1 where id=$2 and deleted_at is null`
	_, err := r.db.ExecContext(ctx, statment, data.Quota, data.Id)
	if err != nil {
		return err
	}
	return nil
}
func (r *JobRepositoryDb) SetExpiredDate(ctx context.Context, data dto.Expiry) error {
	statment := `update jobs set expired_date=$1 where id=$2 and deleted_at is null`
	_, err := r.db.ExecContext(ctx, statment, data.ExpiredDate, data.Id)
	if err != nil {
		return err
	}
	return nil
}
