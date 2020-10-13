package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Car...
type Car struct {
	ID           primitive.ObjectID `bson:"_id"`
	Address      string             `bson:"address"`
	LicensePlate string             `bson:"license_plate"`
	FrameID      string             `bson:"frame_id"`
	EngineID     string             `bson:"engine_id"`
	CarSizeID    primitive.ObjectID `bson:"car_size_id"`
	OwnerID      primitive.ObjectID `bson:"owner_id"`
}
