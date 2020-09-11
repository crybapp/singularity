package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Portal represents an entry in the "Portal" collection. It stores information about VM endpoints and how to connect to them.
type Portal struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	RoomID primitive.ObjectID `bson:"room_id" json:"room_id"`

	Status string `bson:"status" json:"status"`

	RTPAudioPort string `bson:"rtp_audio_port" json:"rtp_audio_port"`
	RTPVideoPort string `bson:"rtp_video_port" json:"rtp_video_port"`

	CreatedAt     primitive.DateTime `bson:"created_at" json:"created_at"`
	LastUpdatedAt primitive.DateTime `bson:"last_updated_at" json:"last_updated_at"`
}
