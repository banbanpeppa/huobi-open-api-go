package rest

import (
	"strings"
)

type Stock struct {
	Amount    float64 `json:"amount"`
	Ts        int64   `json:"ts"`
	Price     float64 `json:"price"`
	Direction string  `json:"direction"`
}
type Data struct {
	Id   int64   `json:"id"`
	Data []Stock `json:"data"`
}

type StockResponse struct {
	Status string `json:"status"`
	Data   []Data `json:"data"`
	Ts     int64  `json:"ts"`
	Ch     string `json:"ch"`
}

func (handler *Handler) GetStock(symbols []string) {
	strRequest := "/market/history/trade"
	go func() {
		for {
			for _, symbol := range symbols {
				params := make(map[string]string)
				params["symbol"] = strings.ToLower(symbol) + "usdt"
				params["size"] = "1"
				handler.processSymbol(params, strRequest, StockResponse{})
			}
			// time.Sleep(500 * time.Microsecond) //睡眠
		}
	}()
}
