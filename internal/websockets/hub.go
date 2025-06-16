package websockets

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Name string
	Conn *websocket.Conn
}

type Hub struct {
	Clients   map[string]*Client
	Broadcast chan string
	mu        sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:   make(map[string]*Client),
		Broadcast: make(chan string),
	}
}

func (h *Hub) HandleMessages() {
	for {
		msg := <-h.Broadcast
		h.mu.Lock()
		for name, client := range h.Clients {
			err := client.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("WebSocket error: ", err)
				client.Conn.Close()
				delete(h.Clients, name)
			}
		}
		h.mu.Unlock()
	}
}
