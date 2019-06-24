package rest

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

func (handler *Handler) SubscribeSpotTrade(symbols []string) {
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

func (handler *Handler) SubscribeSpotDepth(symbols []string, depth DepthRequestDepth, depthType DepthStep) {
	strRequest := "/market/depth"
	go func() {
		for {
			for _, symbol := range symbols {
				params := make(map[string]string)
				params["symbol"] = strings.ToLower(symbol) + "usdt"
				if depth != DEPTH_DEFAULT {
					params["depth"] = strconv.Itoa(int(depth))
				}
				params["type"] = string(depthType)
				handler.processSymbol(params, strRequest, DepthResponse{})
				// time.Sleep(500 * time.Millisecond)
			}
		}
	}()
}

func (handler *Handler) GetSpotDepth(symbol string, depth DepthRequestDepth, depthType DepthStep) (*DepthResponse, error) {
	strRequest := "/market/depth"
	params := make(map[string]string)
	params["symbol"] = strings.ToLower(symbol) + "usdt"
	if depth != DEPTH_DEFAULT {
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

func (handler *Handler) GetSpotDepths(symbols []string, depth DepthRequestDepth, depthType DepthStep) ([]*DepthResponse, error) {
	depths := make([]*DepthResponse, 0)
	for _, symbol := range symbols {
		depth, err := handler.GetSpotDepth(symbol, depth, depthType)
		if err != nil {
			return nil, err
		} else {
			depths = append(depths, depth)
		}
	}
	return depths, nil
}
