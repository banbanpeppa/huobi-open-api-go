package rest

import (
	"time"
)

type Index struct {
	Symbol     string  `json:"symbol"`
	IndexPrice float64 `json:"index_price"`
	IndexTs    int64   `json:"index_ts"`
}

type IndexResponse struct {
	Status string  `json:"status"`
	Data   []Index `json:"data"`
	Ts     int64   `json:"ts"`
}

/**
 * 获取合约指数
 *
 * @param symbol
 *            ["BTC","ETH"...]
 * @return
 */
func (handler *Handler) GetFutureContractIndex(symbols []string) {
	strRequest := "/api/v1/contract_index"
	go func() {
		for {
			for _, symbol := range symbols {
				params := make(map[string]string)
				params["symbol"] = symbol
				go handler.processSymbol(params, strRequest, IndexResponse{})
			}
			time.Sleep(time.Second) //睡眠
		}
	}()
}

func (handler *Handler) GetFutureMarketTrade(symbols []string) {
	strRequest := "/market/trade"
	suffixs := []string{"_CQ", "_CW", "_NW"}
	go func() {
		for {
			for _, symbol := range symbols {
				params := make(map[string]string)
				for _, suf := range suffixs {
					params["symbol"] = symbol + suf
					handler.processSymbol(params, strRequest, nil)
				}
			}
			// time.Sleep(time.Second) //睡眠
		}
	}()
}