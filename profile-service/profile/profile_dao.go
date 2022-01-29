package profile

import (
	"context"
	"log"

	"github.com/user-service/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "profile"

type ProfileDocument struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Address string             `bson:"address"`
	Phone   string             `bson:"phone"`
	UserID  primitive.ObjectID `bson:"userId"`
}

func getCollection() *mongo.Collection {
	db := db.GetDatabase()
	return db.Collection(COLLECTION_NAME)
}

func createProfile(address string, phone string, uid string) (ProfileDocument, error) {

	coll := getCollection()

	oid, err := primitive.ObjectIDFromHex(uid)

	if err != nil {
		return ProfileDocument{}, err
	}

	log.Printf("About to create profile with userId %s, address %s, phone %s", oid.Hex(), address, phone)

	res, err := coll.InsertOne(context.TODO(), bson.D{{Key: "address", Value: address}, {Key: "phone", Value: phone}, {Key: "userId", Value: oid}})

	if err != nil {
		return ProfileDocument{}, err
	}

	return ProfileDocument{ID: res.InsertedID.(primitive.ObjectID), Address: address, Phone: phone, UserID: oid}, nil
}

func updateProfile(uid string, p ProfileDocument) (ProfileDocument, error) {
	coll := getCollection()
	var newProfile ProfileDocument
	oid, err := primitive.ObjectIDFromHex(uid)

	if err != nil {
		return ProfileDocument{}, err
	}

	log.Printf("About to update profile with id: %s", oid.Hex())

	err = coll.FindOneAndUpdate(context.TODO(), bson.D{{Key: "_id", Value: oid}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "address", Value: p.Address},
			{Key: "phone", Value: p.Phone}}}}).Decode(&newProfile)

	if err != nil {
		return ProfileDocument{}, err
	}

	return newProfile, nil
}
