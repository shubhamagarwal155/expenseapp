package users

import (
	"github.com/gin-gonic/gin"
	"go.mod/app/domain"
	"go.mod/app/dtos"
)

type Service struct {
	repo domain.UserRepository
}

func (s *Service) LoginUser(ctx *gin.Context) (res *dtos.Result, err error) {

	var dto *dtos.LoginUserDTO = &dtos.LoginUserDTO{}
	if err = ctx.ShouldBindJSON(&dto); err == nil {
		res, err = s.repo.LoginUser(dto)
	}
	return res, err
}

func (s *Service) RegisterUser(ctx *gin.Context) (res *dtos.Result, err error) {
	var dto *dtos.RegisterUserDTO = &dtos.RegisterUserDTO{}

	if err = ctx.ShouldBindJSON(&dto); err == nil {
		res, err = s.repo.RegisterUser(dto)
	}

	return res, err
}
