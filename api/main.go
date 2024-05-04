package main

import (
	"canto/api/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var myEnv map[string]string
	myEnv, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	route.NewRouter(r, myEnv)

}
