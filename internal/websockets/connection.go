package websockets

import (
	"log"
	"net/http"

	"github.com/Sohamsk/messaging/internal/service"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func (h *Hub) Connect(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrader Error: ", err)
		return
	}

	defer conn.Close()

	h.Clients[conn] = true
	service.SendOldMessages(conn)

	for {
		if _, _, err := conn.NextReader(); err != nil {
			delete(h.Clients, conn)
			break
		}
	}
}
