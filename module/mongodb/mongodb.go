package mongodb

import (
	"context"
	"insurancesvc/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

// Connect ...
func Connect() {

	// Load dotenv for database connection
	env := config.GetENV()

	// Get connection
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(env.DatabaseConnection))
	if err != nil {
		log.Fatal(err)
	}

	// Set data
	db = client.Database(env.DatabaseName)
}

// CarCol ...
func CarCol() *mongo.Collection {
	return db.Collection("Cars")
}

// OrganizationCol ...
func OrganizationCol() *mongo.Collection {
	return db.Collection("Organizations")
}

// OwnerCol ...
func OwnerCol() *mongo.Collection {
	return db.Collection("Owners")
}

// UserCol ...
func UserCol() *mongo.Collection {
	return db.Collection("Users")
}

// CarSizeCol ...
func CarSizeCol() *mongo.Collection {
	return db.Collection("CarSizes")
}

// CompulsoryInsuranceCol ...
func CompulsoryInsuranceCol() *mongo.Collection {
	return db.Collection("CompulsoryInsurances")
}

// VoluntaryInsuranceCol ...
func VoluntaryInsuranceCol() *mongo.Collection {
	return db.Collection("VoluntaryInsurances")
}

// AccidentInsuranceCol ...
func AccidentInsuranceCol() *mongo.Collection {
	return db.Collection("AccidentInsurances")
}

// OrderedInsuranceCol ...
func OrderedInsuranceCol() *mongo.Collection {
	return db.Collection("OrderedInsurances")
}

// OrderCol ...
func OrderCol() *mongo.Collection {
	return db.Collection("Orders")
}

// EmployeeCol ...
func EmployeeCol() *mongo.Collection {
	return db.Collection("Employees")
}
