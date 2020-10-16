package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// CarResponse...
type CarResponse struct {
	ID           primitive.ObjectID `json:"_id"`
	Owner        string             `json:"owner"`
	Address      string             `json:"address"`
	LicensePlate string             `json:"license_plate"`
	FrameID      string             `json:"frame_id"`
	EngineID     string             `json:"engine_id"`
	CarSize      CarSizeResponse    `json:"car_size"`
}
