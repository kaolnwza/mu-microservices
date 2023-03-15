package entity

type Hub struct {
	Rooms      map[string]*WSRoom
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
	ActiveRoom chan *WSRoom
}

func NewHub() *Hub {
	// roomConstr := NewRoom()

	return &Hub{
		Rooms:      make(map[string]*WSRoom),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
		ActiveRoom: make(chan *WSRoom),
	}
}
