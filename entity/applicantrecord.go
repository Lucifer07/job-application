package entity

import "time"

type ApplicantRecord struct {
	Id        *int       `json:"id"`
	UserId    *int       `json:"user_id"`
	JobId     *int       `json:"job_id"`
	Status    *string    `json:"status"`
	CreatedAt *time.Time `json:"applied_at"`
}
