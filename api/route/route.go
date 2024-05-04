package route

import (
	"canto/api/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, myEnv map[string]string) {
	r.GET("/eth/uniswap", handler.GetUniswapPoolData(myEnv))
	r.GET("/eth/aave", handler.GetAaveData(myEnv))
	r.GET("/canto/lending", handler.GetCantoLending(myEnv))
	r.GET("/canto/dex", handler.GetCantoDex(myEnv))
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
