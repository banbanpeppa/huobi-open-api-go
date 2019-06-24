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

type DepthStep string

const (
	DEPTH_STEP0 DepthStep = "step0"
	DEPTH_STEP1 DepthStep = "step1"
	DEPTH_STEP2 DepthStep = "step2"
	DEPTH_STEP3 DepthStep = "step3"
	DEPTH_STEP4 DepthStep = "step4"
	DEPTH_STEP5 DepthStep = "step5"
)

type DepthRequestDepth int32

const (
	DEPTH_FIVE    DepthRequestDepth = 5
	DEPTH_TEN     DepthRequestDepth = 10
	DEPTH_TWENTY  DepthRequestDepth = 20
	DEPTH_DEFAULT DepthRequestDepth = 150
)

type FutureSymbolType string

const (
	CW FutureSymbolType = "_CW"
	CQ FutureSymbolType = "_CQ"
	NW FutureSymbolType = "_NW"
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
