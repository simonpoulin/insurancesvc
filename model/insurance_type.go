package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsuranceType...
type InsuranceType struct {
	Name      string             `bson:"name"`
	CarSizeID primitive.ObjectID `bson:"car_size_id"`
}

// InsuranceTypeBSON...
type InsuranceTypeBSON struct {
	Name      string             `bson:"name"`
	CarSizeID primitive.ObjectID `bson:"car_size_id"`
}

// ConvertToBSON ...
func (payload InsuranceTypePayload) ConvertToBSON() (bson InsuranceTypeBSON) {
	bson = InsuranceTypeBSON{
		Name:      payload.Name,
		CarSizeID: payload.CarSizeObjectID,
	}
	return
}
