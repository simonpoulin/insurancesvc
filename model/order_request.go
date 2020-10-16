package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderCreatePayload...
type OrderCreatePayload struct {
	UserID               string      `json:"user_id"`
	CarID                string      `json:"car_id"`
	OrganizationID       string      `json:"organization_id"`
	Receiver             string      `json:"receiver"`
	Address              string      `json:"address"`
	Phone                string      `json:"phone"`
	Email                string      `json:"email"`
	CreatedAt            time.Time   `json:"created_at"`
	Date                 time.Time   `json:"date"`
	Note                 string      `json:"note"`
	Insurances           []Insurance `json:"insurances"`
	Status               string      `json:"status"`
	UserObjectID         primitive.ObjectID
	CarObjectID          primitive.ObjectID
	OrganizationObjectID primitive.ObjectID
}

// OrderUpdatePayload...
type OrderUpdatePayload struct {
	UserID               string      `json:"user_id"`
	CarID                string      `json:"car_id"`
	OrganizationID       string      `json:"organization_id"`
	Receiver             string      `json:"receiver"`
	Address              string      `json:"address"`
	Phone                string      `json:"phone"`
	Email                string      `json:"email"`
	Date                 time.Time   `json:"date"`
	Note                 string      `json:"note"`
	Insurances           []Insurance `json:"insurances"`
	Status               string      `json:"status"`
	UserObjectID         primitive.ObjectID
	CarObjectID          primitive.ObjectID
	OrganizationObjectID primitive.ObjectID
}
