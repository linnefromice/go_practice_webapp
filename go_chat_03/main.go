package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", &templateHandler{filename: "/chat.html"})
	chatroom := newRoom()
	http.Handle("/room", chatroom)
	go chatroom.run()

	log.Println("Lanuching chat server.")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Failure on launching chat server.: ", err)
	}
}