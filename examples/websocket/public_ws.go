package main

import (
	"fmt"

	bybit "github.com/buurzx/bybit.go.api"
)

func main() {
	ws, errCh := bybit.NewBybitPublicWebSocket("wss://stream.bybit.com/v5/public/spot", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	})
	_ = ws.Connect([]string{"orderbook.1.BTCUSDT"})
	select {
	case <-errCh:
		ws.Disconnect()
		ws.Connect([]string{"orderbook.1.BTCUSDT"})
	}
}
