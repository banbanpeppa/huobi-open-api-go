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
也可以通过`api`的形式调用websocket的订阅
```
var TICKER_ALL = []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
p := websocket.NewDefaultParameters()
huobiClient := websocket.NewHuobiWSClient(p) //WS运行太久，外部原因可能断开，支持自动重连

huobiClient.SubscribeFutureMarketTrade(TICKER_ALL)
for obj := range huobiClient.Listen() {
    switch obj.(type) {
    case string:
        fmt.Print(obj)
    case *websocket.TradeDetail:
        abc := obj.(*websocket.TradeDetail)
        fmt.Println(abc.Tick.Id)
    case []byte:
        go func() {
            tradeDetail := &websocket.TradeDetail{}
            err := json.Unmarshal(obj.([]byte), &tradeDetail)
            if err == nil {
                now := time.Now().UTC().Unix()
                if len(tradeDetail.Tick.Data) > 0 && (now-tradeDetail.Tick.Data[0].Ts/1000) < 5 {
                    price := tradeDetail.Tick.Data[0].Price
                    fmt.Println(tradeDetail.Ch+" ", price)
                }
            }
        }()
    }
}
```

#### Huobi Websocket API

[合约Websocket 文档](https://github.com/huobiapi/API_Docs/wiki/WS_api_reference_Derivatives)

### RESTFul
### Basic requests
```
handler := rest.NewDefaultFutureRestHandler()
tickers := []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
handler.SubscribeFutureMarketDepth(tickers, rest.STEP0)
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
```