package services

import "go.mongodb.org/mongo-driver/mongo"

//ServerContext defines our global dependencies.
type ServerContext interface {
	GetMongoDBClient() *mongo.Client
	GetApplicationConfig() Config
}
