package users

import (
	"database/sql"
)

func Wire(db *sql.DB) *Handler {
	userRepository := ProvideRepository(db)
	userService := ProvideService(userRepository)
	userHandler := ProvideHandler(userService)
	return userHandler
}
