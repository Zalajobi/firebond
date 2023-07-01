package handlers

import (
	"encoding/json"
	"firebond/db"
	"firebond/helpers"
	"firebond/models"
	"firebond/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type PriceResponse struct {
	Price float64 `json:"price"`
}

type MultipleRatesResponse struct {
	Fiat  string
	Value float64
}

func GetPrice(c *gin.Context) {
	var fiat, cryptoCurrency string = strings.ToLower(c.Param("fiat")), strings.ToLower(c.Param("cryptocurrency"))

	// Fetch price from CoinGecko API via request
	if fiat == "" || cryptoCurrency == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fiat and cryptocurrency parameters are required"})
		return
	}

	// Call CoinGecko API to get price
	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + cryptoCurrency + "&vs_currencies=" + fiat
	response, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()

	var result map[string]map[string]float64
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	price := result[cryptoCurrency][fiat]

	// Create the response payload
	priceResponse := PriceResponse{
		Price: price,
	}

	c.JSON(http.StatusOK, priceResponse)
}

func GetCryptocurrencyRates(c *gin.Context) {
	var crypto string = strings.ToLower(c.Param("cryptocurrency"))
	rates := []MultipleRatesResponse{}

	for _, fiat := range utils.SupportedFiatCurrencies {
		rate, err := helpers.GetCryptocurrencyRate(crypto, fiat)

		if err != nil {
			log.Printf("Failed to get %s/%s rate: %v", crypto, fiat, err)
		} else {
			rates = append(rates, MultipleRatesResponse{
				Fiat:  fiat,
				Value: rate,
			})
		}
	}

	c.JSON(http.StatusOK, rates)
}

func GetAllCryptoCurrencyRate(c *gin.Context) {
	rates := []MultipleRatesResponse{}
	for _, crypto := range utils.SupportedCryptocurrencies {
		for _, fiat := range utils.SupportedFiatCurrencies {
			crypto = strings.ToLower(crypto)
			fiat = strings.ToLower(fiat)

			rate, err := helpers.GetCryptocurrencyRate(crypto, fiat)
			if err != nil {
				log.Printf("Failed to get %s/%s rate: %v", crypto, fiat, err)
			} else {
				rates = append(rates, MultipleRatesResponse{
					Fiat:  fmt.Sprintf("%s/%s", crypto, fiat),
					Value: rate,
				})
			}
		}
	}

	c.JSON(http.StatusOK, rates)
}

func GetCryptoCurrencyHistory(c *gin.Context) {
	var fiat, crypto string = strings.ToLower(c.Param("fiat")), strings.ToLower(c.Param("cryptocurrency"))
	var rates []models.Rates

	dbConn := db.GetDB()
	err := dbConn.Where("created_at >= ? AND name = ?", time.Now().Add(-24*time.Hour), fmt.Sprintf("%s/%s", crypto, fiat)).Find(&rates).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rates)
}
