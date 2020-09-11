package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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
