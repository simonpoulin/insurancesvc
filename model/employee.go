package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Employee...
type Employee struct {
	ID       primitive.ObjectID
	Name     string
	Address  string
	Phone    string
	Password string
	Email    string
	Active   bool
}
