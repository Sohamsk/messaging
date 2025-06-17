package handlers

import (
	"net/http"
	"text/template"

	"github.com/Sohamsk/messaging/internal/service/sessions"
	"github.com/google/uuid"
)

var templ = template.Must(template.ParseFiles("./views/message.page.html"))

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("username")
	if name == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	sessionId := uuid.New()
	sessions.Create(sessionId.String(), name)
	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: sessionId.String(),

		HttpOnly: true,
		// Secure: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	templ.Execute(w, nil)
}
