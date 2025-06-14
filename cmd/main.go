package main

import (
	"log"
	"net/http"

	"github.com/Sohamsk/messaging/internal/handlers"
	"github.com/Sohamsk/messaging/internal/websockets"
)

var PORT = "8000"

func main() {
	fs := http.FileServer(http.Dir("./static"))
	hub := websockets.NewHub()

	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("POST /send", handlers.HandleSend(hub))
	http.HandleFunc("/ws", hub.Connect)
	http.HandleFunc("/", handlers.ServeHome)
	log.Println("Starting server at PORT " + PORT)

	go hub.HandleMessages()

	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
