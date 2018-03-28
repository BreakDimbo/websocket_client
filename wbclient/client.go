package wbclient

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	addr   string
	isStop bool
	wsConn *websocket.Conn
	mu     *sync.Mutex
}

func New(addr string) *Client {
	return &Client{
		addr:   addr,
		isStop: false,
		mu:     &sync.Mutex{},
	}
}

func (c *Client) Connect() {
	go c.run()
}

func (c *Client) WsConn() *websocket.Conn {
	return c.wsConn
}

func (c Client) SendMessage(msg []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.wsConn.WriteMessage(websocket.TextMessage, msg)
}

func (c *Client) run() {
	u := url.URL{Scheme: "ws", Host: c.addr, Path: "/ws"}
	fmt.Printf("connecting to %s\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
		return
	}
	defer conn.Close()

	c.wsConn = conn

	// 接收消息
	go func() {
		defer conn.Close()
		for !c.isStop {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("[ERROR] read: %s\n", err)
				return
			}

			if string(message) == "pong" {
				continue
			}

			fmt.Printf("recv message: %s\n", string(message))
		}
	}()

	// 心跳和停止消息
	ticker := time.NewTicker(50 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := c.SendMessage([]byte("ping"))
			if err != nil {
				fmt.Printf("[ERROR] write: %s\n", err)
				return
			}
		}
	}
}
