package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// EmployeeResponse...
type EmployeeResponse struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Address  string             `json:"address"`
	Phone    string             `json:"phone"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
	Active   bool               `json:"active"`
}
