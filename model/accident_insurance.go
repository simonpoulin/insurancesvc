package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// AccidentInsurance...
type AccidentInsurance struct {
	ID    primitive.ObjectID
	Price int
	Value int
}
