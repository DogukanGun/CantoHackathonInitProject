package handler

import (
	"canto/canto/lending"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCantoLending(myEnv map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		client, err := ethclient.Dial(myEnv["CANTO_URI"])
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "Canto network now unreachable",
				"data":  nil,
			})
			return
		}
		pools := lending.GetLendingPoolsApy(client)
		c.JSON(http.StatusOK, map[string]interface{}{
			"data": pools,
		})
	}
}
