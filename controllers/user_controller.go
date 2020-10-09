package controllers

import (
	"crybapp/singularity/models"
	"crybapp/singularity/services"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//UserController is a wrapper around Controller to prevent method contamination.
type UserController struct {
	*Controller
}

//RegisterUserController registers the userController and all it's routes to the router.
func RegisterUserController(baseRoute string, context services.ServerContext, router *httprouter.Router) {
	userController := UserController{
		NewController(baseRoute, context, router),
	}

	userController.GET("me", userController.getSelf)
	userController.POST("register", userController.register)

}

func (userController UserController) getSelf(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
}

func (userController UserController) register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

	newUser := models.User{}

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
	}

	newJSON, err := json.Marshal(newUser)
	w.Write(newJSON)
}
