package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/rafaelsouzaribeiro/web-chat-with-websocket-using-a-map-variable-in-go/web/templates"
)

func (server *Server) serveChat(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").Parse(templates.Chat)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		WebSocketURL string
	}{
		WebSocketURL: fmt.Sprintf("ws://%s:%d%s", server.host, server.port, server.pattern),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
