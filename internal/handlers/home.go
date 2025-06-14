package handlers

import (
	"log"
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	isHTMX := r.Header.Get("HX-Request") != ""
	isAjax := r.Header.Get("X-Requested-With") == "XMLHttpRequest"

	if r.URL.Path != "/" {
		log.Println(r.URL, " not found")
		if isHTMX || isAjax {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "./views/404.html")
		return
	}
	http.ServeFile(w, r, "./views/index.html")
}
