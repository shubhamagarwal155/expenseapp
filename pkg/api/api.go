package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"go.mod/app/routes"
)

type ApiServer struct {
	db     *sql.DB
	router *gin.Engine
}

func NewApiServer(db *sql.DB) *ApiServer {
	server := &ApiServer{
		db: db,
	}
	router := gin.Default()
	server.router = router

	return server
}

func (server *ApiServer) Run(addr string) error {

	routes.RegisterUserRoutes(server.router, server.db)

	return server.router.Run(addr)
}
