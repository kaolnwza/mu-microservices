package port

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type WSRoom interface {
	JoinRoom(conn *websocket.Conn, roomUUID string, clientUUID string)
	RoomUpgrader(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
	ActiveRoom(roomUUID string)
}
