package main

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
const mongoURI = "mongodb://localhost:27017/" + dbName

// Connect to the MongoDB database
// Connect to the MongoDB database
func Connect() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
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
func FindUser(email string) (User, error) {
	log.Info("Finding user with email: ", email)
	var user User
	err := mg.Db.Collection("users").FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		log.Errorf("Error finding user: %v", err)
		return user, err
	}
	return user, nil
}

// GetGiftIdsForUser Get all gifts from a user
func GetGiftIdsForUser(email string) ([]string, error) {
	log.Info("Getting gifts for user with email: ", email)
	var user User
	err := mg.Db.Collection("users").FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		log.Errorf("Error getting gifts for user: %v", err)
		return nil, err
	}
	return user.Gifts, nil
}

// InsertGiftForUser Insert a gift for a user
func InsertGiftForUser(email string, giftId string) error {
	_, err := mg.Db.Collection("users").UpdateOne(context.TODO(), bson.D{{"email", email}}, bson.M{"$push": bson.M{"gifts": giftId}})
	if err != nil {
		return err
	}
	return nil
}

// DeleteGiftForUser Delete a gift for a user
func DeleteGiftForUser(email string, giftId string) error {
	_, err := mg.Db.Collection("users").UpdateOne(context.TODO(), bson.D{{"email", email}}, bson.M{"$pull": bson.M{"gifts": giftId}})
	if err != nil {
		return err
	}
	return nil
}

func ListAllGifts() ([]Gift, error) {
	var gifts []Gift
	cursor, err := mg.Db.Collection("gifts").Find(context.TODO(), bson.M{})
	if err != nil {
		return gifts, err
	}
	if err = cursor.All(context.TODO(), &gifts); err != nil {
		return gifts, err
	}
	return gifts, nil
}

func GetGiftById(id string) (Gift, error) {
	// Convert string id to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Error converting id to ObjectID: %v", err)
		return Gift{}, err
	}

	var gift Gift
	err = mg.Db.Collection("gifts").FindOne(context.TODO(), bson.D{{"_id", objectID}}).Decode(&gift)
	if err != nil {
		log.Errorf("Error finding gift: %v", err)
		return gift, err
	}
	return gift, nil
}

func InsertGift(gift Gift) (string, error) {
	id, err := mg.Db.Collection("gifts").InsertOne(context.TODO(), gift)
	if err != nil {
		return "", err
	}
	return id.InsertedID.(primitive.ObjectID).Hex(), nil
}

func GetGiftByIds(ids []string) ([]Gift, error) {
	log.Info("Getting gifts by ids: ", ids)
	var objectIds []primitive.ObjectID
	for _, id := range ids {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Errorf("Error converting id to ObjectID: %v", err)
			return nil, err
		}
		objectIds = append(objectIds, objectID)

	}
	var gifts []Gift
	cursor, err := mg.Db.Collection("gifts").Find(context.TODO(), bson.M{"_id": bson.M{"$in": objectIds}})
	if err != nil {
		return gifts, err
	}
	if err = cursor.All(context.TODO(), &gifts); err != nil {
		return gifts, err
	}
	return gifts, nil
}
