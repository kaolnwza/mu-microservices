package ws

import (
	entity "github.com/kaolnwza/muniverse/chat/internal/application/core/entities"
)

type wsHub struct {
	*entity.Hub
	cli *entity.Client
}

func NewWSHub() *wsHub {
	hubConstr := entity.NewHub()

	// hubConstr.Rooms["704cd95b-2ae9-474a-90e8-cbd544b8a498"] = &entity.WSRoom{
	// 	UUID:    "704cd95b-2ae9-474a-90e8-cbd544b8a498",
	// 	Name:    "yesped",
	// 	Clients: make(map[string]*entity.Client),
	// }

	return &wsHub{Hub: hubConstr}
}

func (h *wsClient) Run() {
	for {
		select {
		case cl := <-h.hub.ActiveRoom:
			if _, ok := h.hub.Rooms[cl.UUID]; !ok {
				h.hub.Rooms[cl.UUID] = cl
			}

		case cl := <-h.hub.Register:
			if _, ok := h.hub.Rooms[cl.RoomUUID]; ok {
				r := h.hub.Rooms[cl.RoomUUID]
				if _, ok := r.Clients[cl.UUID]; !ok {
					r.Clients[cl.UUID] = cl
				}
			}

		case cl := <-h.hub.Unregister:
			if _, ok := h.hub.Rooms[cl.RoomUUID]; ok {
				if _, ok := h.hub.Rooms[cl.RoomUUID].Clients[cl.UUID]; ok {
					if len(h.hub.Rooms[cl.RoomUUID].Clients) != 0 {
						h.hub.Broadcast <- &entity.Message{
							Type:     entity.MESSAGE_TYPE_NOTICE,
							Text:     "chat_left",
							RoomUUID: cl.RoomUUID,
							UserUUID: cl.UUID,
						}
					}

					delete(h.hub.Rooms[cl.RoomUUID].Clients, cl.UUID)
					close(cl.Message)
				}
			}

		case m := <-h.hub.Broadcast:
			if _, ok := h.hub.Rooms[m.RoomUUID]; ok {
				for _, cl := range h.hub.Rooms[m.RoomUUID].Clients {
					cl.Message <- m
				}
			}
		}
	}
}
