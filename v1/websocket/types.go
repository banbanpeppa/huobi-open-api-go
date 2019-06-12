package websocket

type Data struct {
	Ts        int     `json:"ts"`
	Price     float64 `json:"price"`
	Direction string  `json:"direction"`
}

type Tick struct {
	Id   int    `json:"id"`
	Ts   int    `json:"ts"`
	Data []Data `json:"data"`
}

type TradeDetail struct {
	Ts   int    `json:"ts"`
	Ch   string `json:"ch"`
	Tick Tick   `json:"tick"`
}
