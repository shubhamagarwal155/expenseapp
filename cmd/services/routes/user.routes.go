package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouteHandler struct {
}

func NewHandler() *RouteHandler {
	return &RouteHandler{}
}

func (routeHandler *RouteHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", routeHandler.handleLogin).Methods("POST")
	router.HandleFunc("/registerUser", routeHandler.handleRegisterUser).Methods("POST")
}

func (routeHandler *RouteHandler) handleRegisterUser(writer http.ResponseWriter, request *http.Request) {

}

func (routerHandler *RouteHandler) handleLogin(writer http.ResponseWriter, request *http.Request) {

}
