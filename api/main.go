package main

import (
	"canto/api/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	route.NewRouter(r)

}
