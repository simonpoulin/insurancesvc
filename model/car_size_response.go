package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// CarSizeResponse...
type CarSizeResponse struct {
	ID          primitive.ObjectID `json:"_id"`
	Description string             `json:"desc"`
	MinSeat     int                `json:"min_seat"`
	MaxSeat     int                `json:"max_seat"`
}
