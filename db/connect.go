package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbName    string
	dbPass    string
	dbUri     string
	dbConnStr string
)

func init() {
	dbName = os.Getenv("RUMAIR_DATABASE")
	dbPass = os.Getenv("RUMAIR_DATABASE_PASSWORD")

	if dbName == "" || dbPass == "" {
		log.Fatalf("RUMAIR_DATABASE environment variable must be the name of the Cosmos DB database and RUMAIR_DATABASE_PASSWORD must be the primary password for that database.")
	}

	dbUri = fmt.Sprintf("%s.documents.azure.com:10255", dbName)
	dbConnStr = fmt.Sprintf("mongodb://%s:%s@%s/%s?ssl=true", dbName, dbPass, dbUri, dbName)
}

func CheckMongoDbConnection() {
	var (
		dbClient *mongo.Client
		err      error
	)
	if dbClient, err = mongo.NewClient(options.Client().ApplyURI(dbConnStr)); err != nil {
		log.Fatalf("Can't create mongodb client, go error %v\n", err)
	}

	// Set a 7s timeout and ensure that it is called
	context, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Establish a connection to Cosmos DB
	if err = dbClient.Connect(context); err != nil {
		log.Fatalf("Can't connect to mongodb on %s, go error %v\n", dbName, err)
	}

	// Retrieve a reference to the package collection in Cosmos DB
	collection := dbClient.Database(dbName).Collection("package")

	// Write a single document into Cosmos DB
	var insertResult *mongo.InsertOneResult
	if insertResult, err = collection.InsertOne(context,
		bson.M{
			"FullName":      "react",
			"Description":   "A framework for building native apps with React.",
			"ForksCount":    11392,
			"StarsCount":    48794,
			"LastUpdatedBy": "shergin",
		},
	); err != nil {
		log.Fatalf("Failed to insert record: %v\n", err)
	}

	id := insertResult.InsertedID
	fmt.Println("Inserted record id:", id)

	// Update document
	var updateResult *mongo.UpdateResult
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"fullname": "react-native",
		},
	}
	if updateResult, err = collection.UpdateOne(context, filter, update); err != nil {
		log.Fatalf("Error updating record %v\n", err)
	}
	fmt.Println("Updated this many records:", updateResult.ModifiedCount)

	// Delete document by id
	var deleteResult *mongo.DeleteResult
	if deleteResult, err = collection.DeleteOne(context, filter); err != nil {
		log.Fatalf("Error deleting record: %v\n", err)
	}
	fmt.Println("Deleted this many records:", deleteResult.DeletedCount)

	log.Println("Database was SUCESSFULLY contacted via MongoDB API. DB name is :", dbName)
}
