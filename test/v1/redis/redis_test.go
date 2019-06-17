package test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/banbanpeppa/huobi-open-api-go/utils"
	simplejson "github.com/bitly/go-simplejson"
)

var (
	REGEX_HUOBI_FUTURE_PATTERN = "future:([a-z]+):HUOBI:([A-Z]+)"
	REGEX_HUOBI_INDEX_PATTERN  = "index:HUOBI:([A-Z]+)"
	REGEX_HUOBI_STOCK_PATTERN  = "stock:HUOBI:([A-Z]+)"
	REGEX_HUOBI_PATTERN        = "(.*):HUOBI:(.*)"
	REGEX_OKEX_PATTERN         = "(.*):OKEX:(.*)"
	REGEX_BFX_PATTERN          = "(.*):BFX:(.*)"
)

func TestRedisKeys(t *testing.T) {
	redis := utils.RedisUtils{}
	redis.ConnectByUrl("redis://47.75.66.40:6380")
	defer redis.Close()
	keys, err := redis.Keys()
	if err != nil {
		return
	}
	fmt.Println("一共的key：", len(keys))
	for {
		count := 0
		for _, k := range keys {
			match, _ := regexp.MatchString(REGEX_HUOBI_PATTERN, k)
			if match {
				response, err := redis.Get(k)
				if err == nil {
					j, err := simplejson.NewJson([]byte(response))
					if err == nil {
						ts := j.Get("ts").MustInt64()
						now := time.Now().UTC().Unix()
						if now-ts > 3 {
							count++
						}
					}
				}
			}

		}
		fmt.Println("超时的数目：", count)
	}

}
