package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	}
}
