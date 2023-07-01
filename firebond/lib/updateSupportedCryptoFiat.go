package lib

import (
	"encoding/json"
	"firebond/models"
	"firebond/utils"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func UpdateRateData(db *gorm.DB) {
	for _, crypto := range utils.SupportedCryptocurrencies {
		for _, fiat := range utils.SupportedFiatCurrencies {
			rate, err := GetCryptocurrencyRate(crypto, fiat)
			if err != nil {
				log.Printf("Failed to get Bitcoin/USD rate: %v", err)
			} else {
				rateModel := &models.Rates{
					Name:      fmt.Sprintf("%s/%s", crypto, fiat),
					CreatedAt: time.Now(),
					Crypto:    crypto,
					Fiat:      fiat,
					Value:     rate,
				}
				result := db.Create(&rateModel)
				if result.Error != nil {
					log.Printf("Failed to create rate for Bitcoin/USD: %v", result.Error)
				}
			}
		}
	}
}

func GetCryptocurrencyRate(crypto, fiat string) (float64, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + crypto + "&vs_currencies=" + fiat
	response, err := http.Get(url)
	if err != nil {
		return 0, nil
	}
	defer response.Body.Close()

	var result map[string]map[string]float64
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return 0, nil
	}

	return result[crypto][fiat], nil
}
