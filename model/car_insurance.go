package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CarInsurance...
type CarInsurance struct {
	ID          primitive.ObjectID
	CarSizeID   primitive.ObjectID
	Price       int
	Description string
	Type        string
}
