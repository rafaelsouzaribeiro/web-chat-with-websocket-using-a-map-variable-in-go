package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/web-chat-with-websocket-using-a-map-variable-in-go/internal/usecase/dto"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	mu.Lock()
	for _, v := range buffer {
		conn.WriteJSON(v)
	}
	mu.Unlock()

	defer func() {
		username := getUsernameByConnection(conn)

		mu.Lock()
		delete(messageExists, conn)
		mu.Unlock()

		if username != "" {
			disconnectionMessage := dto.Payload{
				Username: "<strong>info</strong>",
				Message:  fmt.Sprintf("User <strong>%s</strong> disconnected", username),
			}

			mu.Lock()
			buffer = append(buffer, disconnectionMessage)
			mu.Unlock()

			broadcast <- disconnectionMessage

			deleteUserByUserName(username, true)
			conn.Close()
		}
	}()

	for {
		var msgs dto.Payload
		err := conn.ReadJSON(&msgs)
		if err != nil {
			break
		}

		if !verifyExistsUser(msgs.Username, conn) {
			if verifyCon(conn, &messageExists) {
				systemMessage := dto.Payload{
					Username: "<strong>info</strong>",
					Message:  fmt.Sprintf("User already exists: <strong>%s</strong>", msgs.Username),
				}

				deleteUserByConn(conn, false)
				conn.WriteJSON(systemMessage)
			}
			continue
		}

		if msgs.Type == "message" {
			mu.Lock()
			msgs.Username = fmt.Sprintf("<strong>%s</strong>", msgs.Username)
			buffer = append(buffer, msgs)
			mu.Unlock()
			broadcast <- msgs
		} else {
			systemMessage := dto.Payload{
				Username: "<strong>info</strong>",
				Message:  fmt.Sprintf("User <strong>%s</strong> connected", msgs.Username),
			}

			mu.Lock()
			id := uuid.New().String()
			users[id] = User{
				conn:     conn,
				username: msgs.Username,
				id:       id,
			}

			buffer = append(buffer, systemMessage)
			mu.Unlock()

			broadcast <- systemMessage
		}
	}
}
