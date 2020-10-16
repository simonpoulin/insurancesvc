package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CarSize ...
type CarSize struct {
	ID          primitive.ObjectID `bson:"_id"`
	Description string             `bson:"desc"`
	MinSeat     int                `bson:"min_seat"`
	MaxSeat     int                `bson:"max_seat"`
}

// CarSizeCreateBSON ...
type CarSizeCreateBSON struct {
	ID          primitive.ObjectID `bson:"_id"`
	Description string             `bson:"desc"`
	MinSeat     int                `bson:"min_seat"`
	MaxSeat     int                `bson:"max_seat"`
}

// CarSizeUpdateBSON ...
type CarSizeUpdateBSON struct {
	Description string `bson:"desc"`
	MinSeat     int    `bson:"min_seat"`
	MaxSeat     int    `bson:"max_seat"`
}

// ConvertToCreateBSON ...
func (payload CarSizeCreatePayload) ConvertToCreateBSON() (bson CarSizeCreateBSON) {
	bson = CarSizeCreateBSON{
		ID:          primitive.NewObjectID(),
		Description: payload.Description,
		MinSeat:     payload.MinSeat,
		MaxSeat:     payload.MinSeat,
	}
	return
}

// ConvertToUpdateBSON ...
func (payload CarSizeUpdatePayload) ConvertToUpdateBSON() (bson CarSizeUpdateBSON) {
	bson = CarSizeUpdateBSON(payload)
	return
}
