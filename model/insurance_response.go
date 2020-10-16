package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// InsuranceResponse...
type InsuranceResponse struct {
	ID          primitive.ObjectID    `json:"_id"`
	CarSize     CarSizeResponse       `json:"car_size"`
	Price       int                   `json:"price"`
	Description string                `json:"desc"`
	Type        InsuranceTypeResponse `json:"type"`
}
