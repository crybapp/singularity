package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Room represents an entry in the "Rooms" collection. It is used to setup A/V and Chat contexts for users to join.
type Room struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Name   string `bson:"name" json:"name"`
	Invite string `bson:"invite" json:"invite"`

	OwnerID      primitive.ObjectID `bson:"owner_id" json:"owner_id"`
	ControllerID primitive.ObjectID `bson:"controller_id" json:"controller_id"`

	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	EndedAt   primitive.DateTime `bson:"ended_at" json:"ended_at"`
}
