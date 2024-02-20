package dto

type Expiry struct {
	Id          int    `form:"id" json:"id" binding:"required"`
	ExpiredDate string `form:"expired_date" json:"expired_date" binding:"required"`
}
