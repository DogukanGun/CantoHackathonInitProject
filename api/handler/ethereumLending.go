package handler

import (
	"canto/ethereum/aave"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAaveData() gin.HandlerFunc {
	return func(c *gin.Context) {
		pools := aave.GetApyForAavePools(nil)
		c.JSON(http.StatusOK, map[string]interface{}{
			"data": pools,
		})
	}
}
