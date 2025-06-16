package handlers

import (
	"net/http"

	"github.com/Sohamsk/messaging/internal/service"
	"github.com/Sohamsk/messaging/internal/websockets"
)

// TODO: make a format that each message can use that contains the sender name and stuff
func HandleSend(h *websockets.Hub) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		msg := r.FormValue("message")
		service.SaveMessages(msg)

		h.Broadcast <- msg
	}
}
