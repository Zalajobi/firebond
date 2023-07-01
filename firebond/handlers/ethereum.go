package handlers

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
)

type BalanceResponse struct {
	Address string
	Balance *big.Float
}

// GetAccountEthereumBalance Get Ethereum Balance Handler
func GetAccountEthereumBalance(c *gin.Context) {
	address := strings.ToLower(c.Param("address"))
	apiKey := os.Getenv("DEV_INFURIA_API_KEY")

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + apiKey)
	if err != nil {
		log.Fatal(err)
	}

	ethAddress := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), ethAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the balance to Ether
	ethBalance := new(big.Float)
	ethBalance.SetString(balance.String())
	ethBalance = new(big.Float).Quo(ethBalance, big.NewFloat(1e18))

	balanceResponse := BalanceResponse{
		Address: ethAddress.Hex(),
		Balance: ethBalance,
	}

	c.JSON(http.StatusOK, balanceResponse)
}
