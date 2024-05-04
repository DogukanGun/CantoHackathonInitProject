package handler

import (
	"canto/ethereum/aave"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAaveData(myEnv map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		client, err := ethclient.Dial(myEnv["ETH_URI"])
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "Ethereum network now unreachable",
				"data":  nil,
			})
			return
		}
		pools := aave.GetApyForAavePools(client)
		c.JSON(http.StatusOK, map[string]interface{}{
			"data": pools,
		})
	}
}
