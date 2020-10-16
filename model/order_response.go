package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderResponse...
type OrderResponse struct {
	ID             primitive.ObjectID   `json:"_id"`
	UserID         UserResponse         `json:"user"`
	CarID          CarResponse          `json:"car"`
	OrganizationID OrganizationResponse `json:"organization"`
	Receiver       string               `json:"receiver"`
	Address        string               `json:"address"`
	Phone          string               `json:"phone"`
	Email          string               `json:"email"`
	CreatedAt      time.Time            `json:"created_at"`
	Date           time.Time            `json:"date"`
	Note           string               `json:"note"`
	Insurances     []InsuranceResponse  `json:"insurances"`
	Status         string               `json:"status"`
}
