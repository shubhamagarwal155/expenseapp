package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"go.mod/app/users"
)

func RegisterUserRoutes(router *gin.Engine, db *sql.DB) {
	//get user handler
	userHandler := users.Wire(db)

	userRouterGroup := router.Group("/user")
	userRouterGroup.POST("/login", userHandler.Login)
}
