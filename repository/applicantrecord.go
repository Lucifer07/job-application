package repository

import (
	"context"
	"database/sql"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
)

type ApplicantRecordRepositoryDb struct {
	db *sql.DB
}
type ApplicantRecordRepository interface {
	AppliedJob(ctx context.Context, record dto.AppliedReq) (*int, error)
	FindAppliedJob(ctx context.Context, jobId int, userId int) (*dto.ApplicantRecord, error)
}

func NewAppicantRecordRepository(db *sql.DB) *ApplicantRecordRepositoryDb {
	return &ApplicantRecordRepositoryDb{
		db: db,
	}
}
func (r *ApplicantRecordRepositoryDb) AppliedJob(ctx context.Context, record dto.AppliedReq) (*int, error) {
	var id int
	db := util.GetQueryRunner(ctx, r.db)
	statment := `insert into aplication_record(user_id,job_id)values($1,$2) returning id;`
	err := db.QueryRowContext(ctx, statment, record.UserId, record.JobId).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
func (r *ApplicantRecordRepositoryDb) FindAppliedJob(ctx context.Context, jobId int, userId int) (*dto.ApplicantRecord, error) {
	db := util.GetQueryRunner(ctx, r.db)
	u := dto.UserResponse{}
	j := entity.Job{}
	a := dto.ApplicantRecord{}
	statment := `select a.id,a.user_id,u.username,a.job_id,j.name,j.quota,j.expired_date,a.status,a.created_at from aplication_record a join users u on a.user_id=u.id join jobs j on j.id=a.job_id where a.job_id=$1 and a.user_id=$2;`
	err := db.QueryRowContext(ctx, statment, jobId, userId).Scan(&a.Id, &u.Id, &u.Username, &j.Id, &j.Name, &j.Quota, &j.ExpiredDate, &a.Status, &a.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	a.JobId = j.Id
	a.Job = j
	a.UserId = u.Id
	a.User = u
	return &a, nil
}
