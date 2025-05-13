package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	var client *mongo.Client
	var err error

	for i := 0; i < 5; i++ {
		client, err = mongo.Connect(ctx, clientOpts)
		if err == nil && client.Ping(ctx, nil) == nil {
			fmt.Println("Successfully connected to MongoDB!")
			return client
		}
		fmt.Printf("Retrying MongoDB connection... (%d/5)\n", i+1)
		time.Sleep(3 * time.Second)
	}

	panic(fmt.Errorf("failed to connect to MongoDB: %w", err))
}

func CreatePatientsCollection(db *mongo.Database) error {
	validator := bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"first_name"},
			"properties": bson.M{
				"first_name": bson.M{
					"bsonType":    "string",
					"description": "must be a string and is required",
				},
				"last_name": bson.M{
					"bsonType":    "string",
					"description": "must be a string if present",
				},
			},
		},
	}

	opts := options.CreateCollection().SetValidator(validator)
	err := db.CreateCollection(context.TODO(), "patients", opts)
	if err != nil && !mongo.IsDuplicateKeyError(err) {
		return fmt.Errorf("failed to create patients collection: %w", err)
	}
	return nil
}
