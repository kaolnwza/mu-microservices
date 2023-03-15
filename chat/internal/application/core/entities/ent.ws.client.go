package entity

import "github.com/gorilla/websocket"

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	UUID     string `json:"uuid"`
	RoomUUID string `json:"room_uuid"`
}

type MessageType string

const (
	MESSAGE_TYPE_TYPING MessageType = "typing"
	MESSAGE_TYPE_TEXT   MessageType = "text"
	MESSAGE_TYPE_NOTICE MessageType = "notice"
)

var MessageTypeMapping = map[string]MessageType{
	"typing": MESSAGE_TYPE_TYPING,
	"text":   MESSAGE_TYPE_TEXT,
	"notice": MESSAGE_TYPE_NOTICE,
}

type Message struct {
	Type     MessageType `json:"type"`
	Text     string      `json:"text"`
	UserUUID string      `json:"user_uuid"`
	RoomUUID string      `json:"-"`
}

type MessageRequest struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// type Content struct {
// 	Type string `json:"type"`
// 	Text string `json:"text"`
// }
