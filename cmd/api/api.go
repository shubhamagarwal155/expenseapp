package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mod/cmd/services/routes"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (server *ApiServer) Run() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api").Subrouter()

	userRouteService := routes.NewHandler()
	userRouteService.RegisterRoutes(subRouter)

	log.Fatal(http.ListenAndServe(server.addr, router))
}
