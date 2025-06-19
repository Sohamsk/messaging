package handlers

import (
	"net/http"
	"time"

	"github.com/Sohamsk/messaging/internal/models"
	"github.com/Sohamsk/messaging/internal/service"
	"github.com/Sohamsk/messaging/internal/service/sessions"
	"github.com/Sohamsk/messaging/internal/websockets"
)

// TODO: make a format that each message can use that contains the sender name and stuff
func HandleSend(h *websockets.Hub) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		username, err := sessions.GetUserName(cookie.Value)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		if _, exists := h.Clients[username]; !exists {
			http.Error(w, "WebSocket disconnected!!! cannot procceed", http.StatusConflict)
			return
		}

		msg := r.FormValue("message")
		mess := models.Message{
			Username:  username,
			Content:   msg,
			Timestamp: time.Now(),
		}
		service.SaveMessages(mess)

		h.Broadcast <- mess
	}
}
