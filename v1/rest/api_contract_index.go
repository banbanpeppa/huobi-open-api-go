package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	simplejson "github.com/bitly/go-simplejson"
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
				go handler.processSymbol(symbol, strRequest)
			}
			time.Sleep(time.Second) //睡眠
		}
	}()
	log.Println("get contrace index start.")
}

func (handler *Handler) processSymbol(symbol string, strRequest string) {
	params := make(map[string]string)
	params["symbol"] = symbol
	responsej, err := ApiKeyGet(params, strRequest, handler.Params)
	if err != nil {
		log.Println("get future contract index err:", err)
	} else {
		simJ, err := simplejson.NewJson([]byte(responsej))
		if err != nil {
			log.Println("json umarchal contract index err:", err)
		}
		status := simJ.Get("status").MustString()
		if status == HTTP_OK {
			indexResponse := IndexResponse{}
			err = json.Unmarshal([]byte(responsej), &indexResponse)
			if err != nil {
				handler.listener <- fmt.Sprint("json unmarchal contract index err:", err)
			} else {
				handler.listener <- &indexResponse
			}
		} else if status == HTTP_ERROR {
			e := Error{}
			err = json.Unmarshal([]byte(responsej), &e)
			if err != nil {
				handler.listener <- fmt.Sprint("json unmarchal contract index err:", err)
			} else {
				handler.listener <- &e
			}
		}
	}
	time.Sleep(500 * time.Millisecond)
}
