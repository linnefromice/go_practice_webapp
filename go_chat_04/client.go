package main

import (
	// "bytes"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)
var (
	newline = []byte{'\n'}
	space = []byte{' '}
)
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	// temporary solution : origin
	CheckOrigin: func(r *http.Request) bool {
        return true
	},
}

// Client クライアント
type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan *Post
}
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		var post *Post
		if err := c.conn.ReadJSON(&post); err == nil {
			t := time.Now()
			layout := "2006-01-02 15:04:05"
			post.Time = t.Format(layout)
			c.hub.broadcast <- post
		} else {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case post, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteJSON(post); err != nil {
				c.conn.Close()
			}
		case <- ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func buildClient(hub *Hub, conn *websocket.Conn) *Client {
	log.Println("Client生成")
	return &Client{
		hub: hub,
		conn: conn,
		send: make(chan *Post),
	}
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := buildClient(hub, conn)
	client.hub.register <- client
	go client.writePump()
	go client.readPump()
}
