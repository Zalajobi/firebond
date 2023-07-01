package stores

import (
	"firebond/db"
	"firebond/models"
	"fmt"
	"time"
)

func GetFiatCryptoExchangeLastDay(crypto, fiat string) ([]models.Rates, error) {
	var rates []models.Rates

	dbConn := db.GetDB()
	err := dbConn.Where("created_at >= ? AND name = ?", time.Now().Add(-24*time.Hour), fmt.Sprintf("%s/%s", crypto, fiat)).Find(&rates).Error

	return rates, err
}
