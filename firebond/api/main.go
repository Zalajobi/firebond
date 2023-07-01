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
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ := db.SetupPostgres()
	router := gin.Default()
	baseUrl := os.Getenv("BASE_URL")

	router.GET(baseUrl+"/rates/:cryptocurrency/:fiat", handlers.GetPrice)
	router.GET(baseUrl+"/rates/:cryptocurrency", handlers.GetCryptocurrencyRates)
	router.GET(baseUrl+"/rates", handlers.GetAllCryptoCurrencyRate)
	router.GET(baseUrl+"/rates/history/:cryptocurrency/:fiat", handlers.GetCryptoCurrencyHistory)
	router.GET(baseUrl+"/ethereum/balance/:address", handlers.GetAccountEthereumBalance)

	go func() {
		for range time.Tick(4 * time.Minute) {
			lib.UpdateRateData(db)
		}
	}()

	log.Fatal(router.Run(os.Getenv("DEV_PORT")))
}
