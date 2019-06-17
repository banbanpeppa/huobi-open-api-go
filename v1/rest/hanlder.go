package rest

import (
	"encoding/json"
	"fmt"
	"log"

	simplejson "github.com/bitly/go-simplejson"
)

const (
	HTTP_OK    = "ok"
	HTTP_ERROR = "error"
)

type Error struct {
	Ts      int    `json:"ts"`
	Status  string `json:"status"`
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

type Handler struct {
	Params   *ApiParameter
	listener chan interface{}
}

func NewDefaultFutureRestHandler() *Handler {
	apiParams := CreateDefaultFutureApiParameter()
	return &Handler{Params: apiParams, listener: make(chan interface{})}
}

func NewDefaultStockRestHandler() *Handler {
	apiParams := CreateDefaultStockApiParameter()
	return &Handler{Params: apiParams, listener: make(chan interface{})}
}

func NewRestHandler(apiParams *ApiParameter) *Handler {
	return &Handler{Params: apiParams, listener: make(chan interface{})}
}

func (handler *Handler) Listen() <-chan interface{} {
	return handler.listener
}

func (handler *Handler) processSymbol(params map[string]string, strRequest string, responseType interface{}) {
	responsej, err := ApiKeyGet(params, strRequest, handler.Params)
	if err != nil {
		log.Println("get response err:", err)
	} else {
		simJ, err := simplejson.NewJson([]byte(responsej))
		if err != nil {
			log.Println("json umarshal response err:", err)
			return
		}
		status := simJ.Get("status").MustString()
		if status == HTTP_OK {
			switch responseType.(type) {
			case IndexResponse:
				indexRes := IndexResponse{}
				err = json.Unmarshal([]byte(responsej), &indexRes)
				if err != nil {
					handler.listener <- fmt.Sprint("json unmarshal err:", err)
				} else {
					handler.listener <- &indexRes
				}
			case StockResponse:
				stockRes := StockResponse{}
				err = json.Unmarshal([]byte(responsej), &stockRes)
				if err != nil {
					handler.listener <- fmt.Sprint("json unmarshal err:", err)
				} else {
					handler.listener <- &stockRes
				}
			default:
				handler.listener <- []byte(responsej)
			}
		} else if status == HTTP_ERROR {
			e := Error{}
			err = json.Unmarshal([]byte(responsej), &e)
			if err != nil {
				handler.listener <- fmt.Sprint("json unmarshal Error{} err:", err)
			} else {
				handler.listener <- &e
			}
		}
	}
}
