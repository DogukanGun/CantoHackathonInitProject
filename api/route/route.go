package route

import (
	"canto/api/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	r.GET("/eth/uniswap", handler.GetUniswapPoolData())
	r.GET("/eth/aave", handler.GetAaveData())
	r.GET("/canto/lending", handler.GetCantoLending())
	r.GET("/canto/dex", handler.GetCantoDex())
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
