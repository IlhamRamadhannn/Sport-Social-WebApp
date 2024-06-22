package user

import "encoding/json"

type UserRequest struct {
	ID          json.Number
	Username    string      `json:"username" binding:"required"`
	Age         json.Number `json:"age" binding:"required,number"`
	Email       string      `json:"email" binding:"required,email"`
	PhoneNumber string      `json:"phone number" binding:"required,e164"`
}
