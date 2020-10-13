package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderedInsurance...
type OrderedInsurance struct {
	ID            primitive.ObjectID
	StartDate     time.Time
	EndDate       time.Time
	InsuranceType interface{}
	Quantity      int
	OrderID       primitive.ObjectID
}
