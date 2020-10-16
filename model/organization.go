package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Organization...
type Organization struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Buyer       string             `bson:"buyer"`
	TaxCode     string             `bson:"tax_code"`
	Address     string             `bson:"address"`
	BankAccount string             `bson:"bank_account"`
}

// OrganizationCreateBSON...
type OrganizationCreateBSON struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Buyer       string             `bson:"buyer"`
	TaxCode     string             `bson:"tax_code"`
	Address     string             `bson:"address"`
	BankAccount string             `bson:"bank_account"`
}

// OrganizationUpdateBSON...
type OrganizationUpdateBSON struct {
	Name        string `bson:"name"`
	Buyer       string `bson:"buyer"`
	TaxCode     string `bson:"tax_code"`
	Address     string `bson:"address"`
	BankAccount string `bson:"bank_account"`
}

// ConvertToCreateBSON ...
func (payload OrganizationCreatePayload) ConvertToCreateBSON() (bson OrganizationCreateBSON) {
	bson = OrganizationCreateBSON{
		ID:          primitive.NewObjectID(),
		Name:        payload.Name,
		Buyer:       payload.Buyer,
		TaxCode:     payload.TaxCode,
		Address:     payload.Address,
		BankAccount: payload.BankAccount,
	}
	return
}

// ConvertToUpdateBSON ...
func (payload OrganizationUpdatePayload) ConvertToUpdateBSON() (bson OrganizationUpdateBSON) {
	bson = OrganizationUpdateBSON(payload)
	return
}
