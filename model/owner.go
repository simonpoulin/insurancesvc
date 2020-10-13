package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Owner...
type Owner struct {
	ID      primitive.ObjectID
	Name    string
	Address string
	Phone   string
	Email   string
}
