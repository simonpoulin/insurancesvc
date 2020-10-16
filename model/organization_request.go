package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// OrganizationCreatePayload...
type OrganizationCreatePayload struct {
	ID          primitive.ObjectID `json:"_id"`
	Name        string             `json:"name"`
	Buyer       string             `json:"buyer"`
	TaxCode     string             `json:"tax_code"`
	Address     string             `json:"address"`
	BankAccount string             `json:"bank_account"`
}

// OrganizationUpdatePayload...
type OrganizationUpdatePayload struct {
	Name        string `json:"name"`
	Buyer       string `json:"buyer"`
	TaxCode     string `json:"tax_code"`
	Address     string `json:"address"`
	BankAccount string `json:"bank_account"`
}
