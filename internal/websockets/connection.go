package websockets

import (
	"log"
	"net/http"

	"github.com/Sohamsk/messaging/internal/service"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func (h *Hub) Connect(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	if username == "" {
		log.Println("Error: empty username attempted by " + r.RemoteAddr)
		http.Error(w, "Need to send username", http.StatusBadRequest)
		return
	}

	h.mu.Lock()
	if _, exists := h.Clients[username]; exists {
		log.Println("Warning: Creating a new connection for " + username + " while killing the old one.")
		h.Clients[username].Conn.Close()
		delete(h.Clients, username)
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrader Error: ", err)
		h.mu.Unlock()
		return
	}

	defer conn.Close()

	h.Clients[username] = &Client{Name: username, Conn: conn}
	service.SendOldMessages(conn)
	h.mu.Unlock()

	log.Println("User: " + username + " connected")

	for {
		if _, _, err := conn.NextReader(); err != nil {
			break
		}
	}
}
