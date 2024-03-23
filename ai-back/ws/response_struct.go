package ws

const (
	CHAT_MESSAGE = "chat_message"
	CHAT_SYSTEM  = "chat_system"
)

type WsReMessage struct {
	Type    string `json:"type"`
	Content any    `json:"content"`
}

type ChatResContent struct {
	UserId  int    `json:"userId"`
	ChatId  int    `json:"chatId"`
	Message string `json:"message"`
}
