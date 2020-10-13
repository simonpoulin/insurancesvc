package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CarSize...
type CarSize struct {
	ID      primitive.ObjectID
	Name    string
	MinSeat int
	MaxSeat int
}
