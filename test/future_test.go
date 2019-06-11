package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/banbanpeppa/huobi-open-api-go/websocket"
)

func TestFutureTicker(t *testing.T) {
	p := websocket.NewDefaultParameters()
	huobiClient := websocket.NewHuobiWSClient(p) //WS运行太久，外部原因可能断开，支持自动重连

	requests := []websocket.Request{}
	for _, ticker := range TICKER_ALL {
		req_cw := websocket.Request{Id: "id7", Sub: "market." + ticker + "_CW.trade.detail"}
		req_nw := websocket.Request{Id: "id7", Sub: "market." + ticker + "_NW.trade.detail"}
		req_cq := websocket.Request{Id: "id7", Sub: "market." + ticker + "_CQ.trade.detail"}
		requests = append(requests, req_cw, req_nw, req_cq)
	}

	huobiClient.Subscribe(requests)
	for obj := range huobiClient.Listen() {
		switch obj.(type) {
		case string:
			fmt.Print(obj)
		case *websocket.TradeDetail:
			abc := obj.(*websocket.TradeDetail)
			fmt.Println(abc.Tick.Id)
		case []byte:
			// fmt.Println(string(obj.([]byte)))
			tradeDetail := &websocket.TradeDetail{}
			err := json.Unmarshal(obj.([]byte), &tradeDetail)
			if err == nil {
				if len(tradeDetail.Tick.Data) > 0 {
					price := tradeDetail.Tick.Data[0].Price
					fmt.Println(tradeDetail.Ch+" ", price)
				}
			}
		}
	}
}

func TestTrimFix(t *testing.T) {
	a := "market.etcusdt.trade.detail"
	b := strings.TrimSuffix(a[7:], "usdt.trade.detail")
	b = strings.ToUpper(b)
	fmt.Println(b)

	str := "market.BTC_CQ.trade.detail"
	ticker := strings.TrimSuffix(str[7:], ".trade.detail") //btc
	fmt.Println(ticker[0 : len(ticker)-3])
}
