package main

import (
	"context"
	"fmt"

	bybit "github.com/buurzx/bybit.go.api"
)

func main() {
	GetCoinInfo()
}

func GetCoinInfo() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"coin": "USDT"}
	assetResult, err := client.NewAssetService(params).GetCoinInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(assetResult))
}
