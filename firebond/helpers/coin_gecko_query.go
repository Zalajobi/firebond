package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

	fmt.Println(result[crypto][fiat])

	return result[crypto][fiat], nil
}
