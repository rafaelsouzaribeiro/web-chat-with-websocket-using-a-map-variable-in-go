package main

import (
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/web-chat-with-websocket-using-a-map-variable-in-go/configs"

	"github.com/rafaelsouzaribeiro/web-chat-with-websocket-using-a-map-variable-in-go/internal/infra/web/websocket/server"
)

func main() {

	Conf, err := configs.LoadConfig("./")

	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(Conf.Port)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	svc := server.NewServer(Conf.HostName, Conf.WsEndPoint, port)
	go svc.ServerWebsocket()
	select {}

}
