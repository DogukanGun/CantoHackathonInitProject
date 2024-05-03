package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUniswapPoolData() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
