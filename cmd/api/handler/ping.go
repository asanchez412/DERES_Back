package handler

import (
	"github.com/andsanchez/DERES_Back/pkg/web"
	"github.com/gin-gonic/gin"
)

func PingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		web.Success(c, 200, "pong")
	}
}
