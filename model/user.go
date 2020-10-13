package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User...
type User struct {
	ID             primitive.ObjectID
	Name           string
	Address        string
	Phone          string
	Email          string
	Sex            bool
	DOB            time.Time
	IdentityNumber string
	Password       string
}
