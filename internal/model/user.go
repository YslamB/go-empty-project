package model

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type UserCrete struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
