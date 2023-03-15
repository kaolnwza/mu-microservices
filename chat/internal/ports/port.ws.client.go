package port

import entity "github.com/kaolnwza/muniverse/chat/internal/application/core/entities"

type WSChatClient interface {
	WritePump(cli *entity.Client)
	ReadPump(cli *entity.Client)

	Run()
}
