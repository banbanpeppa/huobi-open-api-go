package websocket

const (
	ACCESS_KEY string = "*"
	SECRET_KEY string = "*"

	// API请求地址, 不要带最后的/
	MARKET_URL string = "https://api.hbdm.com"
	TRADE_URL  string = "https://api.hbdm.com"

	WS_STOCK_URL        string = "wss://api.huobi.pro/ws"
	WS_FUTURE_URL       string = "wss://www.hbdm.com/ws"
	WS_FUTURE_ORDER_URL string = "ws://api.hbdm.com/notification"

	Local_IP string = "*.*.*.*" //Your Local IP

	//replace with real URLs and HostName
	HOST_NAME string = "api.hbdm.com"
)
