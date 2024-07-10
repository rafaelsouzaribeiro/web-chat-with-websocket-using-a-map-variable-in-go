package main

import (
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/Web-chat-with-WebSocket-using-a-map-variable-in-Go/configs"

	"github.com/spf13/viper"

	"github.com/rafaelsouzaribeiro/Web-chat-with-WebSocket-using-a-map-variable-in-Go/internal/infra/web/websocket/server"
)

func main() {

	viper.AutomaticEnv()
	hostname := viper.GetString("HOST_NAME")
	wsEndpoint := viper.GetString("WS_ENDPOINT")
	portStr := viper.GetString("PORT")

	if hostname == "" {
		Conf, err := configs.LoadConfig("../")

		if err != nil {
			panic(err)
		}

		hostname = Conf.HostName
		wsEndpoint = Conf.WsEndPoint
		portStr = Conf.Port
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	svc := server.NewServer(hostname, wsEndpoint, port)
	go svc.ServerWebsocket()
	select {}

}