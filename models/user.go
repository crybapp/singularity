package models

import (
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

//User represents an entry in the "Users" collection.
type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`

	AccessToken  string `bson:"access_token" json:"access_token"`
	RefreshToken string `bson:"refresh_token" json:"refresh_token"`

	JoinedAt int64 `bson:"joined_at" json:"joined_at"`
}

//NonUniqueUser is the error type for when Username or Email is already in use
type NonUniqueUser struct {
	offendingField string
}

func (err *NonUniqueUser) Error() string {
	return err.offendingField + " is already in use."
}

//InsertUser Inserts the current user struct into the database.
func (user User) InsertUser() (primitive.ObjectID, error) {
	if !packageReady {
		return primitive.NilObjectID, &PackageNotReady{}
	}

	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)

	if FindUserByUsername(user.Username) != nil {
		return primitive.NilObjectID, &NonUniqueUser{offendingField: "Username"}
	}

	if FindUserByEmailAddress(user.Email) != nil {
		return primitive.NilObjectID, &NonUniqueUser{offendingField: "Email Address"}
	}

	user.JoinedAt = time.Now().Unix()

	userDocument, err := bson.Marshal(&user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	objectID, err := dbManager.userCollection.InsertOne(context.TODO(), userDocument)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return objectID.InsertedID.(primitive.ObjectID), nil
}

//FindUserByUsername searches the mongo db for a user with that username
func FindUserByUsername(username string) *User {
	username = strings.ToLower(username)
	foundUser := &User{}
	searchJSON := bson.D{
		bson.E{
			Key:   "username",
			Value: username,
		},
	}

	err := dbManager.userCollection.FindOne(context.TODO(), searchJSON).Decode(foundUser)
	if err != nil {
		return nil
	}

	return foundUser
}

//FindUserByEmailAddress searches the mongo db for a user with that email address
func FindUserByEmailAddress(email string) *User {
	email = strings.ToLower(email)
	var foundUser User
	searchJSON := bson.D{
		bson.E{
			Key:   "email",
			Value: email,
		},
	}

	err := dbManager.userCollection.FindOne(context.TODO(), searchJSON).Decode(&foundUser)
	if err != nil {
		return nil
	}

	return &foundUser
}

//FindUserByObjectID searches the mongo db for a user with that ObjectID
func FindUserByObjectID(ID primitive.ObjectID) *User {
	var foundUser User
	searchJSON := bson.D{
		bson.E{
			Key:   "_id",
			Value: ID,
		},
	}

	err := dbManager.userCollection.FindOne(context.TODO(), searchJSON).Decode(&foundUser)
	if err != nil {
		return nil
	}

	return &foundUser
}
