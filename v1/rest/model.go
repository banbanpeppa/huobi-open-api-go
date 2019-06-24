package rest

type Trade struct {
	Amount    float64 `json:"amount"`
	Ts        int64   `json:"ts"`
	Price     float64 `json:"price"`
	Direction string  `json:"direction"`
}
type Data struct {
	Id   int64   `json:"id"`
	Data []Trade `json:"data"`
}

type TradeResponse struct {
	Status string `json:"status"`
	Data   []Data `json:"data"`
	Ts     int64  `json:"ts"`
	Ch     string `json:"ch"`
}

type DepthRequestType string

const (
	STEP0 DepthRequestType = "step0"
	STEP1 DepthRequestType = "step1"
	STEP2 DepthRequestType = "step2"
	STEP3 DepthRequestType = "step3"
	STEP4 DepthRequestType = "step4"
	STEP5 DepthRequestType = "step5"
)

type Tick struct {
	Bids [][]float64 `json:"bids"`
	Asks [][]float64 `json:"asks"`
}

type DepthResponse struct {
	Status string `json:"status"`
	Tick   Tick   `json:"tick"`
	Ts     int64  `json:"ts"`
	Ch     string `json:"ch"`
}

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
