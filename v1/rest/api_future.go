package rest

import (
	"encoding/json"
	"errors"
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

/**
 * 获取合约指数
 *
 * @param symbol
 *            ["BTC","ETH"...]
 * @return
 */
func (handler *Handler) SubscribeFutureContractIndex(symbols []string) {
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

func (handler *Handler) SubscribeFutureMarketTrade(symbols []string) {
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

func (handler *Handler) SubscribeFutureMarketDepth(symbols []string, depthType DepthStep) {
	strRequest := "/market/depth"
	suffixs := []string{"_CQ", "_CW", "_NW"}
	go func() {
		for {
			for _, symbol := range symbols {
				params := make(map[string]string)
				for _, suf := range suffixs {
					params["symbol"] = symbol + suf
					params["type"] = string(depthType)
					handler.processSymbol(params, strRequest, DepthResponse{})
				}
			}
		}
	}()
}

func (handler *Handler) GetFutureMarketDepth(symbol string, cycle FutureSymbolType, depthType DepthStep) (*DepthResponse, error) {
	strRequest := "/market/depth"
	params := make(map[string]string)
	params["symbol"] = symbol + string(cycle)
	params["type"] = string(depthType)
	responsej, err := ApiKeyGet(params, strRequest, handler.Params)
	if err != nil {
		return nil, err
	} else {
		simJ, err := simplejson.NewJson([]byte(responsej))
		if err != nil {
			return nil, err
		}
		status := simJ.Get("status").MustString()
		if status == HTTP_OK {
			depthRes := DepthResponse{}
			err = json.Unmarshal([]byte(responsej), &depthRes)
			if err != nil {
				return nil, err
			} else {
				return &depthRes, nil
			}
		} else {
			return nil, errors.New(responsej)
		}
	}
}
