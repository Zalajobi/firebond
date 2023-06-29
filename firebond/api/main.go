package main

import (
	"firebond/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	var baseUrl string = os.Getenv("BASE_URL")

	router.GET(baseUrl+"/rates/:cryptocurrency/:fiat", handlers.GetPrice)

	log.Fatal(router.Run(os.Getenv("DEV_PORT")))
}
