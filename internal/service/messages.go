package service

import (
	"sync"

	"github.com/gorilla/websocket"
)

var (
	messages []string
	mu       sync.Mutex
)

func SaveMessages(msg string) {
	mu.Lock()
	messages = append(messages, msg)
	mu.Unlock()
}

func SendOldMessages(conn *websocket.Conn) {
	mu.Lock()
	for _, messsage := range messages {
		conn.WriteMessage(websocket.TextMessage, []byte(messsage))
	}
	mu.Unlock()
}
