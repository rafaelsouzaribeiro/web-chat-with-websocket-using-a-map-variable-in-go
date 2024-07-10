package server

import "fmt"

func handleMessages() {
	for msg := range broadcast {

		for _, user := range users {
			mu.Lock()
			msg.Username = fmt.Sprintf("<strong>%s</strong>", msg.Username)
			err := user.conn.WriteJSON(msg)
			mu.Unlock()

			if err != nil {
				fmt.Println("Error sending message:", err)
				mu.Lock()
				user.conn.Close()
				mu.Unlock()
				deleteUserByUserName(user.username, false)
			}
		}

	}
}
