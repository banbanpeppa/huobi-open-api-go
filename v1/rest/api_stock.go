package rest

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

func (handler *Handler) SubscribeStockTrade(symbols []string) {
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

func (handler *Handler) SubscribeStockDepth(symbols []string, depth DepthRequestDepth, depthType DepthRequestType) {
	strRequest := "/market/depth"
	go func() {
		for {
			for _, symbol := range symbols {
				params := make(map[string]string)
				params["symbol"] = strings.ToLower(symbol) + "usdt"
				if depth != DEFAULT {
					params["depth"] = strconv.Itoa(int(depth))
				}
				params["type"] = string(depthType)
				handler.processSymbol(params, strRequest, DepthResponse{})
				// time.Sleep(500 * time.Millisecond)
			}
		}
	}()
}

func (handler *Handler) GetStockDepth(symbol string, depth DepthRequestDepth, depthType DepthRequestType) (*DepthResponse, error) {
	strRequest := "/market/depth"
	params := make(map[string]string)
	params["symbol"] = strings.ToLower(symbol) + "usdt"
	if depth != DEFAULT {
		params["depth"] = strconv.Itoa(int(depth))
	}
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

func (handler *Handler) GetStockDepths(symbols []string, depth DepthRequestDepth, depthType DepthRequestType) ([]*DepthResponse, error) {
	depths := make([]*DepthResponse, 0)
	for _, symbol := range symbols {
		depth, err := handler.GetStockDepth(symbol, depth, depthType)
		if err != nil {
			return nil, err
		} else {
			depths = append(depths, depth)
		}
	}
	return depths, nil
}
