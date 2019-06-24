package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/banbanpeppa/huobi-open-api-go/v1/websocket"
)

func TestStockTicker(t *testing.T) {
	p := websocket.NewDefaultParameters()
	huobiClient := websocket.NewHuobiWSClient(p) //配置文件须填写本地IP地址，WS运行太久，外部原因可能断开，支持自动重连
	p.URL = websocket.WS_STOCK_URL
	huobiClient = websocket.NewHuobiWSClient(p)
	requests := []websocket.Request{}
	for _, ticker := range TICKER_ALL {
		sub := strings.ToLower(ticker) + "usdt"
		req := websocket.Request{Id: "id2", Sub: "market." + sub + ".trade.detail"}
		requests = append(requests, req)
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
			} else {
				fmt.Println("error:", err)
			}

		}
	}
}
