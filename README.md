# huobi-open-api-go

[![](https://img.shields.io/badge/api-huobi-blue.svg)](https://huobiapi.github.io/docs/spot/v1/cn/)

An implementation of [Huobi-API](https://huobiapi.github.io/docs/spot/v1/cn/).

## Installation
```
go get github.com/banbanpeppa/huobi-open-api-go
```

## Usage

### Websocket
#### Basic requests
```
var TICKER_ALL = []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}

params := websocket.NewDefaultParameters()
huobiClient := websocket.NewHuobiWSClient(params)

//定义想要订阅的ws请求体
requests := []websocket.Request{}
for _, ticker := range TICKER_ALL {
    req_cw := websocket.Request{Id: "id7", Sub: "market." + ticker + "_CW.trade.detail"}
    req_nw := websocket.Request{Id: "id7", Sub: "market." + ticker + "_NW.trade.detail"}
    req_cq := websocket.Request{Id: "id7", Sub: "market." + ticker + "_CQ.trade.detail"}
    requests = append(requests, req_cw, req_nw, req_cq)
}

huobiClient.Subscribe(requests)
for obj := range huobiClient.Listen() {
    switch obj.(type) {
    case []byte:
        tradeDetail := &websocket.TradeDetail{}
        err := json.Unmarshal(obj.([]byte), &tradeDetail)
        if err == nil {
            if len(tradeDetail.Tick.Data) > 0 {
                price := tradeDetail.Tick.Data[0].Price
                fmt.Println(tradeDetail.Ch+": ", price)
            }
        }
    }
    default:
        fmt.Println("other type of obj")
}
```
#### Huobi Websocket API

[合约Websocket 文档](https://github.com/huobiapi/API_Docs/wiki/WS_api_reference_Derivatives)

### RESTFul
### Basic requests
```
import "github.com/banbanpeppa/huobi-open-api-go/v1/rest"

handler := rest.NewDefaultRestHandler()
tickers := []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
handler.GetFutureContractIndex(tickers)
for index := range handler.Listen() {
    switch index.(type) {
    case string:
        fmt.Println(index.(string)) //错误信息打印为string
    case *rest.IndexResponse:
        ir := index.(*rest.IndexResponse)
		fmt.Println(ir.Data[0].IndexPrice)
    case *rest.Error:
        jsonStr, _ := json.Marshal(index)
        fmt.Println(string(jsonStr))
    default:
        fmt.Println("other type of index")
    }
}
```