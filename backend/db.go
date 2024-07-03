package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

// Database settings (insert your own database name and connection URI)
const dbName = "user"
const mongoURI = "mongodb://user:password@localhost:27017/" + dbName

// Connect to the MongoDB database
func Connect() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	db := client.Database(dbName)
	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

// InsertUser Insert a new user into the database
func InsertUser(user User) error {
	_, err := mg.Db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

// FindUser Find a user in the database
func FindUser(username string) (User, error) {
	var user User
	err := mg.Db.Collection("users").FindOne(context.TODO(), User{Username: username}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetGifts Get all gifts from a user
func GetGifts(username string) ([]int64, error) {
	var user User
	err := mg.Db.Collection("users").FindOne(context.TODO(), User{Username: username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user.Gifts, nil
}

func ListAllGifts() ([]Gift, error) {
	var gifts []Gift
	collection := mg.Db.Collection("gifts")

	// Finding all documents in the collection
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error finding documents: %v", err)
	}
	defer cursor.Close(context.Background())

	// Iterating through the cursor to decode each document into a Gift struct
	for cursor.Next(context.Background()) {
		var gift Gift
		if err := cursor.Decode(&gift); err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		gifts = append(gifts, gift)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return gifts, nil
}

func AddGiftForUser(gift Gift) error {
	_, err := mg.Db.Collection("users").InsertOne(context.Background(), gift)
	if err != nil {
		return err
	}
	return nil
}
