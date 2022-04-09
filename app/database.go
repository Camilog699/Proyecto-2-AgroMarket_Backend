package app

import (
	"context"
	// "fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetConnection() *mongo.Database {
	if db != nil {
		return db
	}
	db = initMongo()
	return db
}

func initMongo() *mongo.Database {
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	uri := os.Getenv("DB_URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	return client.Database(database)
}
