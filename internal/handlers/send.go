package handlers

import (
	"log"
	"net/http"

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
		log.Println("Debug: ", username)

		msg := r.FormValue("message")
		service.SaveMessages(msg)

		h.Broadcast <- msg
	}
}
