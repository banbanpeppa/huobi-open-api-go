package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/banbanpeppa/huobi-open-api-go/v1/rest"
)

func TestStock(t *testing.T) {
	handler := rest.NewDefaultSpotRestHandler()
	tickers := []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
	handler.SubscribeSpotTrade(tickers)
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

func TestSubscribeDepth(t *testing.T) {
	handler := rest.NewDefaultSpotRestHandler()
	tickers := []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
	handler.SubscribeSpotDepth(tickers, 5, rest.STEP0)
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

func TestGetDepth(t *testing.T) {
	handler := rest.NewDefaultSpotRestHandler()
	stockDepth, err := handler.GetSpotDepth("BTC", rest.DEFAULT, rest.STEP0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(stockDepth)
	}
}

func TestGetDepths(t *testing.T) {
	handler := rest.NewDefaultSpotRestHandler()
	tickers := []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
	depths, err := handler.GetSpotDepths(tickers, rest.DEFAULT, rest.STEP0)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, depth := range depths {
			fmt.Println(depth)
		}

	}
}
