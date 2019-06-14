package test

import (
	"log"
	"testing"
	"time"

	goex "github.com/nntaoli-project/GoEx"
	huobi "github.com/nntaoli-project/GoEx/huobi"
)

func TestGoEx(t *testing.T) {
	ws := huobi.NewHbdmWs()
	//ws.ProxyUrl("socks5://127.0.0.1:1080")

	for {
		ws.SetCallbacks(func(ticker *goex.FutureTicker) {
			log.Println(ticker.Ticker)
		}, func(depth *goex.Depth) {
			log.Println(">>>>>>>>>>>>>>>")
			log.Println(depth.ContractType, depth.Pair)
			log.Println(depth.BidList)
			log.Println(depth.AskList)
			log.Println("<<<<<<<<<<<<<<")
		}, func(trade *goex.Trade, s string) {
			log.Println(s, trade)
		})

		// t.Log(ws.SubscribeTicker(goex.BTC_USD, goex.QUARTER_CONTRACT))
		// t.Log(ws.SubscribeDepth(goex.BTC_USD, goex.NEXT_WEEK_CONTRACT, 0))
		ws.SubscribeTrade(goex.LTC_USD, goex.THIS_WEEK_CONTRACT)
		ws.SubscribeTrade(goex.BTC_USD, goex.THIS_WEEK_CONTRACT)
		ws.SubscribeTrade(goex.EOS_USD, goex.THIS_WEEK_CONTRACT)
		time.Sleep(3 * time.Second)
	}

}
