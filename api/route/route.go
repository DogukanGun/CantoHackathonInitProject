package route

import (
	"canto/api/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, myEnv map[string]string) {
	r.GET("/eth/dex", handler.GetUniswapPoolData(myEnv))
	r.GET("/eth/lending", handler.GetAaveData(myEnv))
	r.GET("/canto/lending", handler.GetCantoLending(myEnv))
	r.GET("/canto/dex", handler.GetCantoDex(myEnv))
	r.GET("/invest", handler.Invester(myEnv))
	err := r.Run(":3001")
	if err != nil {
		return
	}
}
