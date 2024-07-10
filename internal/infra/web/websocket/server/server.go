package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) ServerWebsocket() {
	router := mux.NewRouter()
	router.HandleFunc("/chat", server.serveChat).Methods("GET")
	router.HandleFunc(server.pattern, handleConnections)

	go handleMessages()

	fmt.Printf("Server started on %s:%d \n", server.host, server.port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", server.host, server.port), router)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
