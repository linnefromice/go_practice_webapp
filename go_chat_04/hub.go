package main

// Hub ハブ
type Hub struct {
	clients map[*Client]bool
	broadcast chan []byte
	register chan *Client // clientsにaddするためのchannel
	unregister chan *Client // clientsから除外するためのchannel
}

func (hub *Hub) run() {
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = true // registerで送られたClientをtrueとともにmapに保存する
		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)
			}
		case message := <-hub.broadcast:
			for client := range hub.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(hub.clients, client)
				}
			}
		}
	}
}

func newHub() *Hub {
	return &Hub{
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),
		clients: make(map[*Client]bool),
	}
}