package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

type chatroom struct {
	forward chan *message
	join chan *client
	leave chan *client
	clients map[*client]bool
}

func (c *chatroom) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("websocketの開設に失敗しました: ", err)
	}

	client := &client{
		socket: socket,
		send: make(chan *message, messageBufferSize),
		room: c,
	}
	c.join <- client
	defer func() {
		c.leave <- client
	}()

	go client.write()
	client.read()
}

func newRoom() *chatroom {
	t := time.Now()
	layout := "2006-01-02 15:04:05"
	fmt.Println("chatroomが生成された。: ", t.Format(layout))
	return &chatroom{
		forward: make(chan *message),
		join: make(chan *client),
		leave: make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (c *chatroom) run() {
	for {
		select {
		case client := <-c.join:
			c.clients[client] = true
			fmt.Printf("クライアントが入室した、現在 %x 人のクライアントが存在\n", len(c.clients))
		case client := <-c.leave:
			delete(c.clients, client)
			fmt.Printf("クライアントが退出した。現在%x人のクライアントが存在\n", len(c.clients))
		case msg := <-c.forward:
			fmt.Println("メッセージを受信しました")
			for target := range c.clients {
				select {
				case target.send <- msg:
					fmt.Println("メッセージの送信に成功しました")
				default:
					fmt.Println("メッセージの送信に失敗しました")
					delete(c.clients, target)
				}
			}
		}
	}
}