package main

import (
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	send chan *message
	room *chatroom
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			t := time.Now()
			layout := "2006-01-02 15:04:05"
			msg.Time = t.Format(layout)
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}