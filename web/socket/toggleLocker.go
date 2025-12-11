package socket

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow any origin — token protects us
		return true
	},
}

func Handler() http.HandlerFunc {
	// Read token from environment variable
	secretToken := os.Getenv("LOCKER_TOKEN")
	if secretToken == "" {
		log.Fatal("LOCKER_TOKEN env variable must be set")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Extract token (GET param or header)
		token := r.URL.Query().Get("token")

		// Validate token
		if token != secretToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Upgrade to WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade error:", err)
			return
		}
		defer conn.Close()

		log.Println("Pi connected!")

		// Echo loop or command loop
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("read error:", err)
				break
			}
			log.Println("Received:", string(msg))
		}
	}
}
