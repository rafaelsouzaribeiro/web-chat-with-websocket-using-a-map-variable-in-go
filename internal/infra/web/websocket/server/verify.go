package server

import "github.com/gorilla/websocket"

func verifyExistsUser(u string, conn *websocket.Conn) bool {
	mu.Lock()
	defer mu.Unlock()
	for _, user := range users {
		if user.conn != conn && u == user.username {
			return false
		}
	}
	return true
}

func verifyCon(s *websocket.Conn, variable *map[*websocket.Conn]bool) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := (*variable)[s]; !exists {
		(*variable)[s] = true
		return true
	}
	return false
}
