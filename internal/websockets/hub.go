package websockets

import (
	"log"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan string
}

func NewHub() *Hub {
	return &Hub{
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan string),
	}
}

func (h *Hub) HandleMessages() {
	for {
		msg := <-h.Broadcast
		for client := range h.Clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("WebSocket error: ", err)
				client.Close()
				delete(h.Clients, client)
			}
		}
	}
}
