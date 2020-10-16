package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserResponse...
type UserResponse struct {
	ID             primitive.ObjectID `json:"_id"`
	Name           string             `json:"name"`
	Address        string             `json:"address"`
	Phone          string             `json:"phone"`
	Email          string             `json:"email"`
	Sex            string             `json:"sex"`
	DOB            time.Time          `json:"dob"`
	IdentityNumber string             `json:"identity_number"`
	Password       string             `json:"password"`
}
