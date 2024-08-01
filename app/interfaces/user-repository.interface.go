package interfaces

import "go.mod/app/dtos"

type UserRepository interface {
	// GetUserByEmail(string)
	// RegisterUser(registerUserDTO *dtos.RegisterUserDTO)
	LoginUser(loginUserDTO *dtos.LoginUserDTO) (*dtos.Result, error)
}
