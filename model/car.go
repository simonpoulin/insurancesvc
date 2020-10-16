package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Car...
type Car struct {
	ID           primitive.ObjectID `bson:"_id"`
	Owner        string             `bson:"owner"`
	Address      string             `bson:"address"`
	LicensePlate string             `bson:"license_plate"`
	FrameID      string             `bson:"frame_id"`
	EngineID     string             `bson:"engine_id"`
	CarSizeID    primitive.ObjectID `bson:"car_size_id"`
}

// CarCreateBSON...
type CarCreateBSON struct {
	ID           primitive.ObjectID `bson:"_id"`
	Owner        string             `bson:"owner"`
	Address      string             `bson:"address"`
	LicensePlate string             `bson:"license_plate"`
	FrameID      string             `bson:"frame_id"`
	EngineID     string             `bson:"engine_id"`
	CarSizeID    primitive.ObjectID `bson:"car_size_id"`
}

// CarUpdateBSON...
type CarUpdateBSON struct {
	Owner        string             `bson:"owner"`
	Address      string             `bson:"address"`
	LicensePlate string             `bson:"license_plate"`
	FrameID      string             `bson:"frame_id"`
	EngineID     string             `bson:"engine_id"`
	CarSizeID    primitive.ObjectID `bson:"car_size_id"`
}

// ConvertToCreateBSON ...
func (payload CarCreatePayload) ConvertToCreateBSON() (bson CarCreateBSON) {
	bson = CarCreateBSON{
		ID:           primitive.NewObjectID(),
		Owner:        payload.Owner,
		Address:      payload.Address,
		LicensePlate: payload.LicensePlate,
		FrameID:      payload.FrameID,
		EngineID:     payload.EngineID,
		CarSizeID:    payload.CarSizeObjectID,
	}
	return
}

// ConvertToUpdateBSON ...
func (payload CarUpdatePayload) ConvertToUpdateBSON() (bson CarUpdateBSON) {
	bson = CarUpdateBSON{
		Owner:        payload.Owner,
		Address:      payload.Address,
		LicensePlate: payload.LicensePlate,
		FrameID:      payload.FrameID,
		EngineID:     payload.EngineID,
		CarSizeID:    payload.CarSizeObjectID,
	}
	return
}
