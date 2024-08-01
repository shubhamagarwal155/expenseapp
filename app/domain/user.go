package domain

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mod/app/dtos"
)

// User Entity
type (
	User struct {
		FirstName  *string
		MiddleName *string
		LastName   *string
		Email      *string
		UserName   *string
		Password   *string
	}

	UserRepository interface {
		LoginUser(loginUserDTO *dtos.LoginUserDTO) (*dtos.Result, error)
		RegisterUser(registerUserDTO *dtos.RegisterUserDTO) (*dtos.Result, error)
	}

	UserService interface {
		LoginUser(ctx *gin.Context) (*dtos.Result, error)
		RegisterUser(ctx *gin.Context) (*dtos.Result, error)
	}

	UserHandler interface {
		Login() http.HandlerFunc
		RegisterUser() http.HandlerFunc
	}
)
