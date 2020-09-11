package controllers

import (
	"crybapp/singularity/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Controller is the type for API Controllers.
type Controller struct {
	context   services.ServerContext
	router    *httprouter.Router
	baseRoute string
}

//NewController creates a new API Controller at the specified route.
func NewController(baseRoute string, context services.ServerContext, router *httprouter.Router) Controller {
	if baseRoute[len(baseRoute)-1] != '/' {
		baseRoute = baseRoute + "/"
	}

	return Controller{
		context:   context,
		router:    router,
		baseRoute: baseRoute,
	}
}

func (controller Controller) handle(method string, route string, handler httprouter.Handle) {
	if route[0] == '/' {
		route = route[1:]
	}

	controller.router.Handle(method, controller.baseRoute+route, handler)
}

//GET registers a get handler on this controller
func (controller Controller) GET(route string, handler httprouter.Handle) {
	controller.handle(http.MethodGet, route, handler)
}

//HEAD registers a head handler on this controller
func (controller Controller) HEAD(route string, handler httprouter.Handle) {
	controller.handle(http.MethodHead, route, handler)
}

//OPTIONS registers an options handler on this controller
func (controller Controller) OPTIONS(route string, handler httprouter.Handle) {
	controller.handle(http.MethodOptions, route, handler)
}

//POST registers a post handler on this controller
func (controller Controller) POST(route string, handler httprouter.Handle) {
	controller.handle(http.MethodPost, route, handler)
}

//PUT registers a put handler on this controller
func (controller Controller) PUT(route string, handler httprouter.Handle) {
	controller.handle(http.MethodPut, route, handler)
}

//PATCH registers a patch handler on this controller
func (controller Controller) PATCH(route string, handler httprouter.Handle) {
	controller.handle(http.MethodPatch, route, handler)
}

//DELETE registers a delete handler on this controller
func (controller Controller) DELETE(route string, handler httprouter.Handle) {
	controller.handle(http.MethodDelete, route, handler)
}
