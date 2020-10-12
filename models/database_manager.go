package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	packageReady bool
	dbManager    *DBManager
)

//DBManager provides a shared storage for all mongoDB Collections
type DBManager struct {
	database          *mongo.Database
	userCollection    *mongo.Collection
	portalCollection  *mongo.Collection
	roomCollection    *mongo.Collection
	messageCollection *mongo.Collection
}

//DBOptions contains configurable database options for the application.
type DBOptions struct {
	databaseName      string
	userCollection    string
	portalCollection  string
	roomCollection    string
	messageCollection string
}

//PackageNotReady is the error type for the package not having it's dependencies set up.
type PackageNotReady struct {
	message string
}

func (err *PackageNotReady) Error() string {
	return "Package has not been initialized. Please call models.Initialize()"
}

//Initialize will setup the database context for use with the models
func Initialize(mongo *mongo.Client, opts ...*DBOptions) {
	dbManager = &DBManager{}
	options := &DBOptions{}

	if len(opts) == 0 {
		options.databaseName = "Cryb"
		options.userCollection = "Users"
		options.portalCollection = "Portals"
		options.roomCollection = "Rooms"
		options.messageCollection = "Messages"
	} else {
		options = opts[0]

		if options.databaseName == "" {
			options.databaseName = "Cryb"
		}
		if options.userCollection == "" {
			options.userCollection = "Users"
		}
		if options.portalCollection == "" {
			options.portalCollection = "Portals"
		}
		if options.roomCollection == "" {
			options.roomCollection = "Rooms"
		}
		if options.messageCollection == "" {
			options.messageCollection = "Messages"
		}
	}

	dbManager.database = mongo.Database(options.databaseName)
	dbManager.userCollection = dbManager.database.Collection(options.userCollection)
	dbManager.portalCollection = dbManager.database.Collection(options.portalCollection)
	dbManager.roomCollection = dbManager.database.Collection(options.roomCollection)
	dbManager.messageCollection = dbManager.database.Collection(options.messageCollection)

	packageReady = true
}

//Options creates a new DBOptions object
func Options() *DBOptions {
	return &DBOptions{}
}

//SetDatabase sets the name of the database to use
func (options *DBOptions) SetDatabase(name string) *DBOptions {
	options.databaseName = name
	return options
}

//SetUserCollection sets the name of the collection to use for Users
func (options *DBOptions) SetUserCollection(name string) *DBOptions {
	options.userCollection = name
	return options
}

//SetPortalCollection sets the name of the collection to use for Portals
func (options *DBOptions) SetPortalCollection(name string) *DBOptions {
	options.portalCollection = name
	return options
}

//SetRoomCollection sets the name of the collection to use for Rooms
func (options *DBOptions) SetRoomCollection(name string) *DBOptions {
	options.roomCollection = name
	return options
}

//SetMessageCollection sets the name of the collection to use for Messages
func (options *DBOptions) SetMessageCollection(name string) *DBOptions {
	options.messageCollection = name
	return options
}
