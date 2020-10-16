package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// InsuranceTypePayload...
type InsuranceTypePayload struct {
	Name            string `json:"name"`
	CarSizeID       string `json:"car_size_id"`
	CarSizeObjectID primitive.ObjectID
}
