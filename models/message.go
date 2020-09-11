package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Message represents an entry in the "Message" collection. It is used to store chat messages for rooms,
type Message struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID primitive.ObjectID `bson:"author_id" json:"author_id"`
	RoomID   primitive.ObjectID `bson:"room_id" json:"room_id"`

	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	Content   string             `bson:"content" json:"content"`
}
