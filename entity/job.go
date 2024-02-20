package entity

import "time"

type Job struct {
	Id          *int       `json:"id"`
	Name        *string    `json:"name" binding:"required"`
	Quota       *int       `json:"quota" binding:"required,min=0"`
	ExpiredDate *time.Time `json:"expired_date" binding:"required"`
}
