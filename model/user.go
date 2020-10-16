package model

import (
	"insurancesvc/util"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User...
type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	Address        string             `bson:"address"`
	Phone          string             `bson:"phone"`
	Email          string             `bson:"email"`
	Sex            string             `bson:"sex"`
	DOB            time.Time          `bson:"dob"`
	IdentityNumber string             `bson:"identity_number"`
	Password       string             `bson:"password"`
}

// UserCreateBSON...
type UserCreateBSON struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	Address        string             `bson:"address"`
	Phone          string             `bson:"phone"`
	Email          string             `bson:"email"`
	Sex            string             `bson:"sex"`
	DOB            time.Time          `bson:"dob"`
	IdentityNumber string             `bson:"identity_number"`
	Password       string             `bson:"password"`
	SearchString   string             `bson:"searchstring"`
}

// UserUpdateBSON...
type UserUpdateBSON struct {
	Name           string    `bson:"name"`
	Address        string    `bson:"address"`
	Email          string    `bson:"email"`
	Sex            string    `bson:"sex"`
	DOB            time.Time `bson:"dob"`
	IdentityNumber string    `bson:"identity_number"`
	Password       string    `bson:"password"`
	SearchString   string    `bson:"searchstring"`
}

// ConvertToCreateBSON ...
func (payload UserCreatePayload) ConvertToCreateBSON() (bson UserCreateBSON) {
	bson = UserCreateBSON{
		ID:             primitive.NewObjectID(),
		Name:           payload.Name,
		Address:        payload.Address,
		Phone:          payload.Phone,
		Email:          payload.Email,
		Sex:            payload.Sex,
		DOB:            payload.DOB,
		IdentityNumber: payload.IdentityNumber,
		Password:       util.Hash(payload.Password),
		SearchString:   util.ConvertToHex(payload.Name),
	}
	return
}

// ConvertToUpdateBSON ...
func (payload UserUpdatePayload) ConvertToUpdateBSON() (bson UserUpdateBSON) {
	bson = UserUpdateBSON{
		Name:           payload.Name,
		Address:        payload.Address,
		Email:          payload.Email,
		Sex:            payload.Sex,
		DOB:            payload.DOB,
		IdentityNumber: payload.IdentityNumber,
		Password:       util.Hash(payload.Password),
		SearchString:   util.ConvertToHex(payload.Name),
	}
	return
}
