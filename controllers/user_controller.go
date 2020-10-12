package controllers

import (
	"crybapp/singularity/authentication"
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
	userController.POST("password_login", userController.authByPassword)

}

func (userController UserController) getSelf(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
}

func (userController UserController) register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write(nil)
		return
	}

	newUser := models.User{}

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write(nil)
		return
	}

	passwordHash := authentication.EncryptPassword(newUser.Password)
	newUser.Password = passwordHash

	objectID, err := newUser.InsertUser()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	createdUser := models.FindUserByObjectID(objectID)
	newJSON, err := json.Marshal(createdUser)
	w.Write(newJSON)
}

func (userController UserController) authByPassword(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write(nil)
		return
	}

	type responseObject struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	form := responseObject{}
	err = json.Unmarshal(body, &form)

	user := models.FindUserByUsername(form.Username)
	if user == nil {
		w.WriteHeader(404)
		w.Write([]byte("User does not exist"))
		return
	}

	isCorrectPassword := authentication.VerifyPassword(form.Password, user.Password)
	if !isCorrectPassword {
		w.WriteHeader(401)
		w.Write([]byte("Incorrect Password"))
		return
	}

	w.Write([]byte("You made it in."))
}
