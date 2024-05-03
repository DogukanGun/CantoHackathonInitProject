package handler

import (
	"canto/canto/lending"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCantoLending() gin.HandlerFunc {
	return func(c *gin.Context) {
		pools := lending.GetLendingPoolsApy(nil)
		c.JSON(http.StatusOK, map[string]interface{}{
			"data": pools,
		})
	}
}
