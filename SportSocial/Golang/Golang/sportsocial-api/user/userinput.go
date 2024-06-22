package user

import (
	"time"
)

type UserInput struct {
	ID          int
	Username    string `json:"username" binding:"required"`
	Age         int    `json:"age" binding:"required,number"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone number" binding:"required,e164"`
	CreatedAt   time.Time
	UpdateAt    time.Time
}
