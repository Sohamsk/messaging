package websockets

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/Sohamsk/messaging/internal/models"
	"github.com/gorilla/websocket"
)

type Client struct {
	Name string
	Conn *websocket.Conn
}

type Hub struct {
	Clients   map[string]*Client
	Broadcast chan models.Message
	mu        sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:   make(map[string]*Client),
		Broadcast: make(chan models.Message),
	}
}

func (h *Hub) HandleMessages() {
	for {
		msg := <-h.Broadcast
		mess, _ := json.Marshal(msg)

		h.mu.Lock()
		for name, client := range h.Clients {
			err := client.Conn.WriteMessage(websocket.TextMessage, []byte(mess))
			if err != nil {
				log.Println("WebSocket error: ", err)
				client.Conn.Close()
				delete(h.Clients, name)
			}
		}
		h.mu.Unlock()
	}
}
