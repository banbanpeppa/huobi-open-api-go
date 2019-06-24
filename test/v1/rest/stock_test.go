package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/banbanpeppa/huobi-open-api-go/v1/rest"
)

func TestStock(t *testing.T) {
	handler := rest.NewDefaultStockRestHandler()
	tickers := []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
	handler.GetStockTrade(tickers)
	for index := range handler.Listen() {
		switch index.(type) {
		case *rest.TradeResponse:
			ir := index.(*rest.TradeResponse)
			now := time.Now().Unix()
			ts := ir.Ts / 1000
			if now-ts < 5 {
				fmt.Println(ir)
			}
		case *rest.Error:
			jsonStr, _ := json.Marshal(index)
			fmt.Println(string(jsonStr))
		}
	}
}

func TestDepth(t *testing.T) {
	handler := rest.NewDefaultStockRestHandler()
	tickers := []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
	handler.GetStockDepth(tickers, 5, rest.STEP0)
	for index := range handler.Listen() {
		switch index.(type) {
		case *rest.DepthResponse:
			ir := index.(*rest.DepthResponse)
			fmt.Println(ir)
		case *rest.Error:
			jsonStr, _ := json.Marshal(index)
			fmt.Println(string(jsonStr))
		}
	}
}
