package dto

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
