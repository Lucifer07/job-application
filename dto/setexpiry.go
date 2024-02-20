package dto

import "time"

type Expiry struct {
	Id          int        `form:"id" json:"id" binding:"required"`
	ExpiredDate *time.Time `form:"expired_date" json:"expired_date" binding:"required"`
}
