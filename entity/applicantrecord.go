package entity

type ApplicantRecord struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	JobId     int       `json:"job_id"`
	Status    string    `json:"status"`
}
