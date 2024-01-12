package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	CreateAt string `json:"create_at"`
}
