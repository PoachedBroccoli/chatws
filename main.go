package main

import (
	"chat_websocket/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	go controller.Hub.Run()
	// TODO:Use Gin or Custom routing
	router := mux.NewRouter()
	router.HandleFunc("/ws", controller.WSHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
