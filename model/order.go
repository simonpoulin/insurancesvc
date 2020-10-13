package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order...
type Order struct {
	ID             primitive.ObjectID
	UserID         primitive.ObjectID
	CarID          primitive.ObjectID
	OrganizationID primitive.ObjectID
	Status         string
}
