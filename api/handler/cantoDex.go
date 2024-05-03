package handler

import (
	cantoDex "canto/canto/dex"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCantoDex() gin.HandlerFunc {
	return func(c *gin.Context) {
		pools := cantoDex.CalculatePrices(nil)
		c.JSON(http.StatusOK, map[string]interface{}{
			"data": pools,
		})
	}
}
