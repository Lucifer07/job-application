package dto

type Quota struct {
	Id    int `form:"id" json:"id" binding:"required"`
	Quota int `form:"quota" json:"quota" binding:"required,min=0"`
}
