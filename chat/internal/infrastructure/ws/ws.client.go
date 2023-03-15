package ws

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	entity "github.com/kaolnwza/muniverse/chat/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/chat/internal/ports"
	log "github.com/kaolnwza/muniverse/chat/lib/logs"
)

type wsClient struct {
	hub *entity.Hub
}

func NewWSChatClient(hub *entity.Hub) port.WSChatClient {
	return &wsClient{hub: hub}
}

func (c *wsClient) WritePump(cli *entity.Client) {
	defer func() {
		cli.Conn.Close()
	}()

	for {
		message, ok := <-cli.Message
		if !ok {
			return
		}

		if err := cli.Conn.WriteJSON(message); err != nil {
			log.Error(err)
		}
	}
}

func (c *wsClient) ReadPump(cli *entity.Client) {
	defer func() {
		c.hub.Unregister <- cli
		cli.Conn.Close()
	}()

	for {
		_, m, err := cli.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error(err)
			}

			break
		}

		msgReq := entity.MessageRequest{}
		if err := json.Unmarshal(m, &msgReq); err == nil {

			msg := &entity.Message{
				Type:     entity.MessageTypeMapping[msgReq.Type],
				Text:     msgReq.Text,
				UserUUID: cli.UUID,
				RoomUUID: cli.RoomUUID,
			}

			// msg := &entity.Message{
			// 	Type:    entity.MESSAGE_TYPE_CONTENT,
			// 	Content: string(m),
			// 	// Content: entity.Content{
			// 	// 	Type: "kuy",
			// 	// 	Text: string(m),
			// 	// },
			// 	UserUUID: cli.UUID,
			// 	RoomUUID: cli.RoomUUID,
			// }

			c.hub.Broadcast <- msg
		}
	}
}
