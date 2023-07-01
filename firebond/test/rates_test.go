package test

import (
	"firebond/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCryptoCurrencyRate(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/firebond/pricing/rates/:cryptocurrency/:fiat", handlers.GetPrice)

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
