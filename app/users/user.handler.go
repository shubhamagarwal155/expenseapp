package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mod/app/domain"
)

type Handler struct {
	svc domain.UserService
}

func (h *Handler) Login(ctx *gin.Context) {
	//will be implemented
	res, err := h.svc.LoginUser(ctx)

	if err == nil {
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusInternalServerError, err)
	}
}

func (h *Handler) RegisterUser(ctx *gin.Context) {
	res, err := h.svc.RegisterUser(ctx)

	if err == nil {
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusInternalServerError, err)
	}
}
