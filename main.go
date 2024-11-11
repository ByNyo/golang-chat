package main

import (
	"golang-chat/pkg/api"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	server := api.NewServer()
	http.Handle("/ws", websocket.Handler(server.HandleWS))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
