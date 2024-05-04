package handler

import (
	"canto/api/helper"
	"canto/api/structs"
	"canto/ethereum/uniswap"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUniswapPoolData(myEnv map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		client, err := ethclient.Dial(myEnv["ETH_URI"])
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "Ethereum network now unreachable",
				"data":  nil,
			})
			return
		}
		pools := helper.GetPools()
		targetPools := make([]structs.Pool, len(pools.Data.LiquidityPools))
		for i, pool := range pools.Data.LiquidityPools {
			targetPool := structs.Pool{}
			price, _, tokenDecimal, err := uniswap.UniswapV3PriceOracle(pool.InputTokens[1].ID, pool.InputTokens[0].ID, client)
			if err != nil {
				fmt.Printf("Error while getting token price from uniswap: %v", err)
				continue
			}
			targetPool.Price = price
			targetPool.PoolAddress = pool.ID
			targetPool.TokenDecimal = tokenDecimal
			targetPool.TargetToken = pool.InputTokens[0]
			targetPools[i] = targetPool
		}
		c.JSON(http.StatusOK, gin.H{
			"data": targetPools,
		})
	}
}
