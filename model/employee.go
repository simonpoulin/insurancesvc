package model

import (
	"insurancesvc/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Employee...
type Employee struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Address  string             `bson:"address"`
	Phone    string             `bson:"phone"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
	Active   bool               `bson:"active"`
}

// EmployeeCreateBSON...
type EmployeeCreateBSON struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Address      string             `bson:"address"`
	Phone        string             `bson:"phone"`
	Password     string             `bson:"password"`
	Email        string             `bson:"email"`
	Active       bool               `bson:"active"`
	SearchString string             `bson:"searchstring"`
}

// EmployeeUpdateBSON...
type EmployeeUpdateBSON struct {
	Name         string `bson:"name"`
	Address      string `bson:"address"`
	Password     string `bson:"password"`
	Email        string `bson:"email"`
	Active       bool   `bson:"active"`
	SearchString string `bson:"searchstring"`
}

// ConvertToCreateBSON ...
func (payload EmployeeCreatePayload) ConvertToCreateBSON() (bson EmployeeCreateBSON) {
	bson = EmployeeCreateBSON{
		ID:           primitive.NewObjectID(),
		Name:         payload.Name,
		Address:      payload.Address,
		Phone:        payload.Phone,
		Password:     util.Hash(payload.Password),
		Email:        payload.Email,
		Active:       payload.Active,
		SearchString: util.ConvertToHex(payload.Name),
	}
	return
}

// ConvertToUpdateBSON ...
func (payload EmployeeUpdatePayload) ConvertToUpdateBSON() (bson EmployeeUpdateBSON) {
	bson = EmployeeUpdateBSON{
		Name:         payload.Name,
		Address:      payload.Address,
		Password:     util.Hash(payload.Password),
		Email:        payload.Email,
		Active:       payload.Active,
		SearchString: util.ConvertToHex(payload.Name),
	}
	return
}
