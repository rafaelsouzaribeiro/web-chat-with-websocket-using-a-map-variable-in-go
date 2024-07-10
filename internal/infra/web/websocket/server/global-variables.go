package server

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/web-chat-with-websocket-using-a-map-variable-in-go/internal/usecase/dto"
)

var broadcast = make(chan dto.Payload)
var users = make(map[string]User)
var messageExists = make(map[*websocket.Conn]bool)
var buffer []dto.Payload
var mu sync.Mutex
