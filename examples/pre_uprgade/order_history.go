package main

import (
	"context"
	"fmt"

	bybit "github.com/buurzx/bybit.go.api"
)

func main() {
	PlaceTrade()
}

func PlaceTrade() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"category": "linear", "settleCoin": "USDT", "limit": 10}
	orderResult, err := client.NewPreUpgradeService(params).GetPreUpgradeOrderHistory(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(orderResult))
}
