package service

import (
	"encoding/json"
	"sync"

	"github.com/Sohamsk/messaging/internal/models"
	"github.com/gorilla/websocket"
)

var (
	messages []models.Message
	mu       sync.Mutex
)

func SaveMessages(msg models.Message) {
	mu.Lock()
	messages = append(messages, msg)
	mu.Unlock()
}

func SendOldMessages(conn *websocket.Conn) {
	mu.Lock()
	for _, messsage := range messages {
		msg, _ := json.Marshal(messsage)
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}
	mu.Unlock()
}
