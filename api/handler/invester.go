package handler

import (
	"canto/api/structs"
	"github.com/gin-gonic/gin"
)

func Invester(myEnv map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data structs.InvesterRequestBody
		if err := c.BindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": "Wrong body type"})
			return
		}
		if data.Chain == "canto" {

		} else if data.Chain == "ethereum" {

		} else {
			c.JSON(400, gin.H{"error": "Unsupported chain"})
			return
		}
	}
}
