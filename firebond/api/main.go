package main

import (
	"firebond/db"
	"firebond/handlers"
	"firebond/lib"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	db, _ := db.SetupPostgres()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	baseUrl := os.Getenv("BASE_URL")

	router.GET(baseUrl+"/rates/:cryptocurrency/:fiat", handlers.GetPrice)
	router.GET(baseUrl+"/rates/:cryptocurrency", handlers.GetCryptocurrencyRates)

	go func() {
		for range time.Tick(4 * time.Minute) {
			lib.UpdateRateData(db)
		}
	}()

	log.Fatal(router.Run(os.Getenv("DEV_PORT")))
}
