package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var PORT = "8000"

var (
	messages []string
	mu       sync.Mutex

	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan string)
	upgrader  = websocket.Upgrader{}
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	isHTMX := r.Header.Get("HX-Request") != ""
	isAjax := r.Header.Get("X-Requested-With") == "XMLHttpRequest"

	if r.URL.Path != "/" {
		log.Println(r.URL, " not found")
		if isHTMX || isAjax {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "./views/404.html")
		return
	}
	http.ServeFile(w, r, "./views/index.html")
}

func handleSend(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	msg := r.FormValue("message")
	mu.Lock()
	messages = append(messages, msg)
	mu.Unlock()

	broadcast <- msg
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrader Error: ", err)
		return
	}

	defer conn.Close()

	clients[conn] = true

	mu.Lock()
	for _, messsage := range messages {
		conn.WriteMessage(websocket.TextMessage, []byte(messsage))
	}
	mu.Unlock()

	for {
		if _, _, err := conn.NextReader(); err != nil {
			delete(clients, conn)
			break
		}
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("WebSocket error: ", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	http.HandleFunc("/send", handleSend)
	http.HandleFunc("/ws", handleWebsocket)
	http.HandleFunc("/", serveHome)
	log.Println("Starting server at PORT " + PORT)

	go handleMessages()

	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
