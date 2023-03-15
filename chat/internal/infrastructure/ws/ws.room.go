package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
	entity "github.com/kaolnwza/muniverse/chat/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/chat/internal/ports"
)

type chatServer struct {
	hub      *entity.Hub
	upgrader websocket.Upgrader
	cli      port.WSChatClient
	// cli2     *entity.Client
}

func NewWSChatServer(h *entity.Hub, upgrader websocket.Upgrader, cli port.WSChatClient) port.WSRoom {
	return &chatServer{
		hub:      h,
		upgrader: upgrader,
		cli:      cli,
		// cli2:     cli2,
	}
}

func (ws *chatServer) RoomUpgrader(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return ws.upgrader.Upgrade(w, r, nil)
}

func (ws *chatServer) ActiveRoom(roomUUID string) {
	room := &entity.WSRoom{
		UUID:    roomUUID,
		Clients: make(map[string]*entity.Client),
	}

	ws.hub.ActiveRoom <- room
}

func (ws *chatServer) JoinRoom(conn *websocket.Conn, roomUUID string, clientUUID string) {
	client := &entity.Client{
		Conn:     conn,
		Message:  make(chan *entity.Message, 10),
		UUID:     clientUUID,
		RoomUUID: roomUUID,
	}

	m := &entity.Message{
		Type:     entity.MESSAGE_TYPE_NOTICE,
		Text:     "chat_join",
		RoomUUID: roomUUID,
		UserUUID: clientUUID,
	}

	ws.hub.Register <- client
	ws.hub.Broadcast <- m

	go ws.cli.WritePump(client)
	go ws.cli.ReadPump(client)

}
