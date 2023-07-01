package lib

import (
	"firebond/helpers"
	"firebond/models"
	"firebond/utils"
	"fmt"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

func UpdateRateData(db *gorm.DB) {
	for _, crypto := range utils.SupportedCryptocurrencies {
		for _, fiat := range utils.SupportedFiatCurrencies {
			crypto = strings.ToLower(crypto)
			fiat = strings.ToLower(fiat)
			rate, err := helpers.GetCryptocurrencyRate(crypto, fiat)
			if err != nil {
				log.Printf("Failed to get %s/%s rate: %v", crypto, fiat, err)
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
		//time.Sleep(8 * time.Second)
	}
}
