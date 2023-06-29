package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PriceResponse struct {
	Price float64 `json:"price"`
}

func GetPrice(c *gin.Context) {
	var fiat, cryptoCurrency string = c.Param("fiat"), c.Param("cryptocurrency")

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
