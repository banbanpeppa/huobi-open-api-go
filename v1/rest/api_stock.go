package rest

import (
	"strconv"
	"strings"
)

func (handler *Handler) GetStockTrade(symbols []string) {
	strRequest := "/market/history/trade"
	go func() {
		for {
			for _, symbol := range symbols {
				params := make(map[string]string)
				params["symbol"] = strings.ToLower(symbol) + "usdt"
				params["size"] = "1"
				handler.processSymbol(params, strRequest, TradeResponse{})
			}
		}
	}()
}

func (handler *Handler) GetStockDepth(symbols []string, depth int, depthType DepthRequestType) {
	strRequest := "/market/depth"
	go func() {
		for {
			for _, symbol := range symbols {
				params := make(map[string]string)
				params["symbol"] = strings.ToLower(symbol) + "usdt"
				if depth >= 5 {
					params["depth"] = strconv.Itoa(depth)
				}
				params["type"] = string(depthType)
				handler.processSymbol(params, strRequest, DepthResponse{})
				// time.Sleep(500 * time.Millisecond)
			}
		}
	}()
}
