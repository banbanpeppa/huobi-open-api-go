package websocket

import "time"

type ClientParameters struct {
	URL                string
	LocalIP            string
	ReconnectTimeout   time.Duration
	CheckTickerTimeout time.Duration
	PangTickerTimeout  time.Duration
	WSDialerTimeout    time.Duration
	WSDialerKeepAlive  time.Duration
	WSMessageTimeout   time.Duration
}

func NewDefaultParameters() *ClientParameters {
	return &ClientParameters{
		URL:                WS_FUTURE_URL,
		LocalIP:            Local_IP,
		ReconnectTimeout:   time.Second * 100,
		CheckTickerTimeout: time.Second * 20,
		PangTickerTimeout:  time.Second * 3,
		WSDialerTimeout:    time.Second * 30,
		WSDialerKeepAlive:  time.Second * 30,
		WSMessageTimeout:   time.Millisecond * 100,
	}
}
