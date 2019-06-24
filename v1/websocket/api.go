package websocket

import "strings"

func (cli *Client) SubscribeFutureMarketTrade(symbols []string) {
	requests := []Request{}
	for _, ticker := range symbols {
		req_cw := Request{Id: "id7", Sub: "market." + ticker + "_CW.trade.detail"}
		req_nw := Request{Id: "id7", Sub: "market." + ticker + "_NW.trade.detail"}
		req_cq := Request{Id: "id7", Sub: "market." + ticker + "_CQ.trade.detail"}
		requests = append(requests, req_cw, req_nw, req_cq)
	}

	cli.Subscribe(requests)
}

func (cli *Client) SubscribeSpotMarketTrade(symbols []string) {
	requests := []Request{}
	for _, ticker := range symbols {
		sub := strings.ToLower(ticker) + "usdt"
		req := Request{Id: "id2", Sub: "market." + sub + ".trade.detail"}
		requests = append(requests, req)
	}

	cli.Subscribe(requests)
}
