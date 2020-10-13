package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Organization...
type Organization struct {
	ID          primitive.ObjectID
	BuyerName   string
	Name        string
	TaxCode     string
	Address     string
	BankAccount string
}
