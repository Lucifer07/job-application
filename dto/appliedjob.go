package dto

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
)

type ApplicantRecord struct {
	Id        int          `json:"id"`
	UserId    int          `json:"user_id"`
	User      UserResponse `json:"user"`
	JobId     int          `json:"job_id"`
	Job       entity.Job   `json:"job"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"applied_at"`
}
type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
