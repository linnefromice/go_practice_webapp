package main

import (
	"log"
)

// Hub ハブ
type Hub struct {
	clients map[string]bool
}

func (h *Hub) addClient(value string) {
	h.clients[value] = true
}
func (h *Hub) removeClient(value string) {
	v, ok := h.clients[value]
	log.Println("argument: ", value)
	log.Println("v: ", v)
	log.Println("ok: ", ok)
	delete(h.clients, value)
}

func main() {
	log.Println("START main")
	hub := new(Hub)
	hub.clients = map[string]bool{"clientFive": true}
	hub.addClient("clientOne")
	hub.addClient("clientTwo")
	hub.removeClient("clientOne")
	hub.removeClient("clientOne")
	hub.removeClient("clientThree")
	hub.removeClient("clientTwo")
	log.Println("FINISH main")
}