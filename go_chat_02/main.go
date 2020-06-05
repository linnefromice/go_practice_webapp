package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocketサーバにつなぎにいくクライアント
var clients = make(map[*websocket.Conn]bool)
// クライアントから受け取るメッセージを格納
var broadcast = make(chan Message)
// WebSocket更新用
var upgrader = websocket.Upgrader{}

// クライアントからはJSON形式で受け取る
type Message struct {
	Message string `json:message`
}

// クライアントのハンドラ
func HandleClients(w http.ResponseWriter, r *http.Request) {
	// goroutineで起動
	go broadcastMessagesToClients()
	// websocketの状態を更新
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket::", err)
	}
	// websocket を閉じる
	defer websocket.Close()
	clients[websocket] = true
	for {
		var message Message
		// メッセージ読み込み
		err := websocket.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred while reading message: %v", err)
			delete(clients, websocket)
			break
		}
		// メッセージを受け取る
		broadcast <- message
	}
}

func broadcastMessagesToClients() {
	for {
		// メッセージ受け取り
		message := <-broadcast
		// クライアントの数だけループ
		for client := range clients {
			// 書き込む
			err := client.WriteJSON(message)
			if err != nil {
				log.Printf("error occurred while writing message to client: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	log.Println("Launching...")
	http.HandleFunc("/", htmlHandler)
	http.HandleFunc("/chat", HandleClients)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting http server::", err)
		return
	}
}