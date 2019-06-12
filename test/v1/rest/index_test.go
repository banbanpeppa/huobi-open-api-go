package test

import (
	"fmt"
	"testing"

	"github.com/banbanpeppa/huobi-open-api-go/v1/rest"
)

func TestIndex(t *testing.T) {
	handler := rest.NewDefaultRestHandler()
	tickers := []string{"BTC", "ETH", "BCH", "EOS", "LTC", "ETC", "BSV", "XRP"}
	handler.GetFutureContractIndex(tickers)
	for index := range handler.Listen() {
		switch index.(type) {
		case string:
			fmt.Println(index.(string)) //错误信息打印为string
		case *rest.IndexResponse:
			ir := index.(*rest.IndexResponse)
			fmt.Println(ir.Data[0].Symbol+":", ir.Data[0].IndexPrice)
		case *rest.Error:
			// jsonStr, _ := json.Marshal(index)
			// fmt.Println(string(jsonStr))
		default:
			fmt.Println("other type of index")
		}
	}
}
