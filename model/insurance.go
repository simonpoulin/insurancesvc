package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insurance...
type Insurance struct {
	ID          primitive.ObjectID `bson:"_id"`
	Price       int                `bson:"price"`
	Value       int                `bson:"value"`
	Description string             `bson:"desc"`
	Type        InsuranceType      `bson:"type"`
}

// InsuranceCreateBSON...
type InsuranceCreateBSON struct {
	ID          primitive.ObjectID `bson:"_id"`
	Price       int                `bson:"price"`
	Value       int                `bson:"value"`
	Description string             `bson:"desc"`
	Type        InsuranceTypeBSON  `bson:"type"`
}

// InsuranceUpdateBSON...
type InsuranceUpdateBSON struct {
	Price       int               `bson:"price"`
	Value       int               `bson:"value"`
	Description string            `bson:"desc"`
	Type        InsuranceTypeBSON `bson:"type"`
}

// ConvertToCreateBSON ...
func (payload InsuranceCreatePayload) ConvertToCreateBSON() (bson InsuranceCreateBSON) {
	bson = InsuranceCreateBSON{
		ID:    primitive.NewObjectID(),
		Price: payload.Price,
		Value: payload.Value,
		Type:  payload.Type.ConvertToBSON(),
	}
	return
}

// ConvertToUpdateBSON ...
func (payload InsuranceUpdatePayload) ConvertToUpdateBSON() (bson InsuranceUpdateBSON) {
	bson = InsuranceUpdateBSON{
		Price: payload.Price,
		Value: payload.Value,
		Type:  payload.Type.ConvertToBSON(),
	}
	return
}
