package entity

type WSRoom struct {
	UUID    string             `json:"uuid"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

func NewWSRoom() *WSRoom {
	return &WSRoom{
		Clients: make(map[string]*Client),
	}
}
