package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mod/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestToken := ctx.Request.Header.Get("token")

		if requestToken != "" {
			claims, msg := utils.ValidateJWTToken(requestToken)

			if msg != "" {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": msg,
				})
			} else {
				ctx.Set("uid", claims.UserId)
				ctx.Next()
			}
		}
	}
}
