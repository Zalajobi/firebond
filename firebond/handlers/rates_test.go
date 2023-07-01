package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPrice(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/firebond/pricing/rates/:cryptocurrency/:fiat", GetPrice)

	req, err := http.NewRequest("GET", "http://localhost:9001/api/v1/firebond/pricing/rates/bitcoin/usd", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
	}
}

func TestGetCryptocurrencyRates(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/firebond/pricing/rates/:cryptocurrency", GetCryptocurrencyRates)

	req, err := http.NewRequest("GET", "http://localhost:9001/api/v1/firebond/pricing/rates/bitcoin", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
	}
}

func GetAllCryptoCurrencyRateTest(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/firebond/pricing/rates", GetAllCryptoCurrencyRate)

	req, err := http.NewRequest("GET", "http://localhost:9001/api/v1/firebond/pricing/rates", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
	}
}

func GetAccountEthereumBalanceTest(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/firebond/ethereum/balance/:address", GetAccountEthereumBalance)

	req, err := http.NewRequest("GET", "http://localhost:9001/api/v1/firebond/pricing/ethereum/balance/0x0d8775f648430679a709e98d2b0cb6250d2887ef", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
	}
}

func GetCryptoCurrencyHistoryTest(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/firebond/rates/history/:cryptocurrency/:fiat", GetCryptoCurrencyHistory)

	req, err := http.NewRequest("GET", "http://localhost:9001/api/v1/firebond/pricing/rates/history/bitcoin/eur", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
	}
}
