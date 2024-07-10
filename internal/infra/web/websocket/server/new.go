package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	host    string
	port    int
	pattern string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type User struct {
	conn     *websocket.Conn
	username string
	id       string
}

func NewServer(host, pattern string, port int) *Server {
	return &Server{
		host:    host,
		port:    port,
		pattern: pattern,
	}
}
