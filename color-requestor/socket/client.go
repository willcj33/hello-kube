package socket

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"strconv"

	"github.com/gorilla/websocket"
	pb "github.com/willcj33/hello-kube/color-requestor/protos"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub        *Hub
	conn       *websocket.Conn
	send       chan []byte
	Config     ClientSettings
	grpcClient pb.ColorGeneratorServiceClient
	mu         sync.Mutex
}

type ClientSettings struct {
	Count    int
	Interval int
	Mode     string
}

type ColorDivTransport struct {
	Which int    `json:"which"`
	Color string `json:"color"`
}

func (c *Client) consume() {
	defer func() {
		c.hub.remove <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v\n", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.sendMessage <- message
	}
}

func (c *Client) emit() {
	ticker := time.NewTicker(pingPeriod)
	log.Printf("interval: %v, count: %v\n", c.Config.Interval, c.Config.Count)
	duration := time.Duration(c.Config.Interval)
	tickerInterval := time.NewTicker(duration * time.Millisecond)
	defer func() {
		ticker.Stop()
		tickerInterval.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case <-tickerInterval.C:
			messages := make(chan ColorDivTransport)
			var wg sync.WaitGroup
			wg.Add(c.Config.Count)

			for i := 0; i < c.Config.Count; i++ {
				go func(divCount int) {
					defer wg.Done()

					color, err := c.grpcClient.GetColor(context.Background(), &pb.Empty{})
					if err != nil {
						log.Println(err)
						return
					}

					transport := ColorDivTransport{
						Which: divCount,
						Color: color.Hex,
					}
					messages <- transport
				}(i)
			}

			go func() {
				for m := range messages {
					if c == nil || c.conn == nil {
						return
					}
					c.conn.SetWriteDeadline(time.Now().Add(writeWait))
					c.mu.Lock()
					err := c.conn.WriteJSON(m)
					c.mu.Unlock()
					if err != nil {
						log.Println(err)
						continue
					}
				}
			}()
			wg.Wait()
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func SocketHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	strInterval := r.URL.Query().Get("interval")
	strCount := r.URL.Query().Get("count")
	mode := r.URL.Query().Get("mode")
	iInterval := int64(5000)
	iCount := int64(1)
	if mode == "" {
		mode = "default"
	}
	if strInterval != "" {
		iInterval, _ = strconv.ParseInt(strInterval, 10, 64)
	}
	if strCount != "" {
		iCount, _ = strconv.ParseInt(strCount, 10, 64)
	}
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
		Config: ClientSettings{
			Count:    int(iCount),
			Interval: int(iInterval),
			Mode:     mode,
		},
		grpcClient: hub.grpcClient,
	}
	client.hub.add <- client
	go client.emit()
	client.consume()
}
