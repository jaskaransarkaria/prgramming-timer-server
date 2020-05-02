package readers

import (
	"github.com/gorilla/websocket"
	"log"
	"github.com/jaskaransarkaria/programming-timer-server/session"
	// "github.com/jaskaransarkaria/programming-timer-server/writers"
)

func NewConnReader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Connection closing:", err)
			session.HandleRemoveUser(conn)
			conn.Close()
			break
			} else {
			log.Println(string(p))
			session.AddUserConnToSession(string(p), conn)
		}
	}
}

func UpdateChannelReader() {
	for {
		recievedUpdate := <- session.UpdateTimerChannel
		session.HandleUpdateSession(recievedUpdate)
	}
}