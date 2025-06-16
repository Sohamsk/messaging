package handlers

import (
	"net/http"
	"text/template"
)

var templ = template.Must(template.ParseFiles("./views/message.page.html"))

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("username")
	data := struct {
		Username string
	}{
		Username: name,
	}

	templ.Execute(w, data)
}
