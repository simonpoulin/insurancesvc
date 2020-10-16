package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// CarCreatePayload...
type CarCreatePayload struct {
	Owner           string `json:"owner"`
	Address         string `json:"address"`
	LicensePlate    string `json:"license_plate"`
	FrameID         string `json:"frame_id"`
	EngineID        string `json:"engine_id"`
	CarSizeID       string `json:"car_size_id"`
	CarSizeObjectID primitive.ObjectID
}

// CarUpdatePayload...
type CarUpdatePayload struct {
	Owner           string `json:"owner"`
	Address         string `json:"address"`
	LicensePlate    string `json:"license_plate"`
	FrameID         string `json:"frame_id"`
	EngineID        string `json:"engine_id"`
	CarSizeID       string `json:"car_size_id"`
	CarSizeObjectID primitive.ObjectID
}
