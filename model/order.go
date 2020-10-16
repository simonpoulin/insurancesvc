package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order...
type Order struct {
	ID             primitive.ObjectID `bson:"_id"`
	UserID         primitive.ObjectID `bson:"user_id"`
	CarID          primitive.ObjectID `bson:"car_id"`
	OrganizationID primitive.ObjectID `bson:"organization_id"`
	Receiver       string             `bson:"receiver"`
	Address        string             `bson:"address"`
	Phone          string             `bson:"phone"`
	Email          string             `bson:"email"`
	CreatedAt      time.Time          `bson:"created_at"`
	Date           time.Time          `bson:"date"`
	Note           string             `bson:"note"`
	Insurances     []Insurance        `bson:"insurances"`
	Status         string             `bson:"status"`
}

// OrderCreateBSON...
type OrderCreateBSON struct {
	ID             primitive.ObjectID `bson:"_id"`
	UserID         primitive.ObjectID `bson:"user_id"`
	CarID          primitive.ObjectID `bson:"car_id"`
	OrganizationID primitive.ObjectID `bson:"organization_id"`
	Receiver       string             `bson:"receiver"`
	Address        string             `bson:"address"`
	Phone          string             `bson:"phone"`
	Email          string             `bson:"email"`
	CreatedAt      time.Time          `bson:"created_at"`
	Date           time.Time          `bson:"date"`
	Note           string             `bson:"note"`
	Insurances     []Insurance        `bson:"insurances"`
	Status         string             `bson:"status"`
}

// OrderUpdateBSON...
type OrderUpdateBSON struct {
	UserID         primitive.ObjectID `bson:"user_id"`
	CarID          primitive.ObjectID `bson:"car_id"`
	OrganizationID primitive.ObjectID `bson:"organization_id"`
	Receiver       string             `bson:"receiver"`
	Address        string             `bson:"address"`
	Phone          string             `bson:"phone"`
	Email          string             `bson:"email"`
	Date           time.Time          `bson:"date"`
	Note           string             `bson:"note"`
	Insurances     []Insurance        `bson:"insurances"`
	Status         string             `bson:"status"`
}

// ConvertToCreateBSON...
func (payload OrderCreatePayload) ConvertToCreateBSON() (bson OrderCreateBSON) {
	bson = OrderCreateBSON{
		ID:             primitive.NewObjectID(),
		UserID:         payload.UserObjectID,
		CarID:          payload.CarObjectID,
		OrganizationID: payload.OrganizationObjectID,
		Receiver:       payload.Receiver,
		Address:        payload.Address,
		Phone:          payload.Phone,
		Email:          payload.Email,
		CreatedAt:      time.Now(),
		Date:           payload.Date,
		Note:           payload.Note,
		Insurances:     payload.Insurances,
		Status:         payload.Status,
	}
	return
}

// ConvertToUpdateBSON...
func (payload OrderUpdatePayload) ConvertToUpdateBSON() (bson OrderUpdateBSON) {
	bson = OrderUpdateBSON{
		UserID:         payload.UserObjectID,
		CarID:          payload.CarObjectID,
		OrganizationID: payload.OrganizationObjectID,
		Receiver:       payload.Receiver,
		Address:        payload.Address,
		Phone:          payload.Phone,
		Email:          payload.Email,
		Date:           payload.Date,
		Note:           payload.Note,
		Insurances:     payload.Insurances,
		Status:         payload.Status,
	}
	return
}
