package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/banbanpeppa/huobi-open-api-go/v1/rest"
	simplejson "github.com/bitly/go-simplejson"
)

func TestIndex(t *testing.T) {
	handler := rest.NewDefaultFutureRestHandler()
	tickers := []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
	handler.GetFutureContractIndex(tickers)
	for index := range handler.Listen() {
		switch index.(type) {
		case string:
			fmt.Println(index.(string)) //错误信息打印为string
		case *rest.IndexResponse:
			ir := index.(*rest.IndexResponse)
			fmt.Println(ir.Data[0].Symbol+":", ir.Data[0].IndexPrice, ": ", ir.Ts/1000)
		case *rest.Error:
			jsonStr, _ := json.Marshal(index)
			fmt.Println(string(jsonStr))
		default:
			fmt.Println("other type of index")
		}
	}
}

func TestTrade(t *testing.T) {
	handler := rest.NewDefaultFutureRestHandler()
	tickers := []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
	handler.GetFutureMarketTrade(tickers)
	for index := range handler.Listen() {
		switch index.(type) {
		case []byte:
			sj, err := simplejson.NewJson(index.([]byte))
			if err == nil {
				ch := sj.Get("ch").MustString()
				status := sj.Get("status").MustString()
				tick, err := sj.Get("tick").Get("data").Array()
				ts := sj.Get("tick").Get("ts").MustInt64()
				if err == nil {
					each_map := tick[0].(map[string]interface{})
					fmt.Println(ch, status, each_map["price"], ts)
				}

			}

		case *rest.IndexResponse:
			ir := index.(*rest.IndexResponse)
			fmt.Println(ir.Data[0].Symbol+":", ir.Data[0].IndexPrice, ": ", ir.Ts/1000)
		case *rest.Error:
			jsonStr, _ := json.Marshal(index)
			fmt.Println(string(jsonStr))
		default:
			fmt.Println("other type of index")
		}
	}
}

func TestFutureDepth(t *testing.T) {
	handler := rest.NewDefaultFutureRestHandler()
	tickers := []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
	handler.GetFutureMarketDepth(tickers, rest.STEP0)
	for index := range handler.Listen() {
		switch index.(type) {
		case *rest.DepthResponse:
			ir := index.(*rest.DepthResponse)
			fmt.Println(ir)
		case *rest.Error:
			jsonStr, _ := json.Marshal(index)
			fmt.Println(string(jsonStr))
		}
	}
}
