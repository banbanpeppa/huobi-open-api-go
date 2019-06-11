package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/banbanpeppa/huobi-open-api-go/utils"

	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
)

/** 订阅websocket的请求格式
{
	"sub": "market.$symbol.kline.$period",
	"id": "id generate by client"
}
*/
type Request struct {
	Id  string `json:"id"`
	Req string `json:"req"`
	Sub string `json:"sub"`
}

type Data struct {
	// Amount int `json:"amount"`
	Ts int `json:"ts"`
	// Id        uint64  `json:"id"`
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

type Client struct {
	Name     string
	Params   *ClientParameters
	Ws       *websocket.Conn
	listener chan interface{}
}

// ws-client的监控对象
type Moniter struct {
	clientNum  int      //huobiws-client数目
	addChan    chan int // 当添加ws-client的时候，使用管道对象管理
	subChan    chan int // 当关闭了ws-client之后，subChan减1
	lastUseSec int
}

var (
	mon           *Moniter
	clientNameNum int
)

func initMoniter() {
	mon = &Moniter{}
	mon.addChan = make(chan int, 1000)
	mon.subChan = make(chan int, 1000)
	go func() {
		for {
			select {
			case <-mon.addChan: //接收管道消息并处理
				mon.clientNum++
			case <-mon.subChan:
				mon.clientNum--
			}
		}
	}()
}

func AddClientNum() {
	mon.addChan <- 1 //发送管道消息
}

func SubClientNum() {
	mon.subChan <- 1
}

func NowSec() int {
	return int(time.Now().UnixNano() / 1000000000)
}

func NewHuobiWSClient(params *ClientParameters) *Client {
	clientNameNum++
	return &Client{Name: cast.ToString(clientNameNum), Params: params, listener: make(chan interface{})}
}

func (cli *Client) Subscribe(requests []Request) {
	initMoniter()
	go cli.sub(requests)
	// time.Sleep(20 * time.Millisecond)
	// cli.reCreateClient()
}

func (cli *Client) Listen() <-chan interface{} {
	return cli.listener
}

func (cli *Client) reCreateClient() {
	go func() {
		time.Sleep(time.Second * 20)
		checkTicker := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-checkTicker.C:
				if mon.clientNum <= 0 {
					NewHuobiWSClient(cli.Params)
				}
			}
		}
	}()
}

func (cli *Client) sub(reqs []Request) {
	AddClientNum()
	dialer := websocket.DefaultDialer
	dialer.NetDial = func(network, addr string) (net.Conn, error) {

		addrs := []string{string(cli.Params.LocalIP)}
		localAddr := &net.TCPAddr{IP: net.ParseIP(addrs[rand.Int()%len(addrs)]), Port: 0}
		d := net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			LocalAddr: localAddr,
			DualStack: true,
		}
		c, err := d.Dial(network, addr)
		return c, err
	}
	c, _, err := dialer.Dial(cli.Params.URL, nil)
	if err != nil {
		log.Println("Dial Erro:", err)
		SubClientNum()
		return
	}

	defer func() {
		c.Close()
		SubClientNum()
	}()

	for _, request := range reqs {
		message, err := json.Marshal(request)
		if err != nil {
			log.Println("json marshal err :", err)
		}
		messgeByte := []byte(message)
		err = c.WriteMessage(websocket.TextMessage, messgeByte)
		if err != nil {
			log.Println("write err :", err)
		}
	}

	go func() {
		pangTicker := time.NewTicker(time.Second * 3)
		for {
			select {
			case <-pangTicker.C:
				message := []byte(fmt.Sprintf("{\"pong\":%d}", time.Now().Unix()))
				err = c.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println("send msg err:", err)
					return
				}
			}
		}
	}()

	for {
		_, zipmsg, err := c.ReadMessage()
		if err != nil {
			log.Println("Read Error : ", err, cli.Name)
			defer c.Close()
			return
		}

		msg, err := utils.ParseGzip(zipmsg)
		if err != nil {
			log.Println("gzip Error : ", err)
		}

		cli.handleMessage(msg)
	}
}

func (cli *Client) handleMessage(msg []byte) {
	cli.listener <- msg
}
