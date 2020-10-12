package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"crybapp/singularity/controllers"
	"crybapp/singularity/models"
	"crybapp/singularity/services"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ApplicationContext is the live context of our application. It contains references to long lived services.
type ApplicationContext struct {
	MongoDbClient     *mongo.Client
	ApplicationConfig services.Config
}

//GetMongoDBClient returns the currently active mongodb client
func (context ApplicationContext) GetMongoDBClient() *mongo.Client {
	return context.MongoDbClient
}

//GetApplicationConfig returns the current configuration.
func (context ApplicationContext) GetApplicationConfig() services.Config {
	return context.ApplicationConfig
}

func main() {
	applicationConfig, err := services.LoadConfigFromFile("../config/singularity.config.json")
	handleErr(err)

	mongoDbClient, err := mongo.NewClient(options.Client().ApplyURI(applicationConfig.DatabaseURI))
	handleErr(err)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = mongoDbClient.Connect(ctx)
	handleErr(err)

	context := ApplicationContext{
		MongoDbClient:     mongoDbClient,
		ApplicationConfig: applicationConfig,
	}

	models.Initialize(context.MongoDbClient)

	router := httprouter.New()
	router.GET("/", index)

	controllers.RegisterUserController("/users/", context, router)

	log.Fatal(http.ListenAndServe(":"+applicationConfig.APIPort, router))
}

func index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
}

func handleErr(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
