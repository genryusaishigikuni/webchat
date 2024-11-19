package models

type Chat struct {
	ChatID      int      `json:"chat_id"`
	ChatName    string   `json:"chat_name"`
	ChatMembers []string `json:"chat_members"`
}

type RegisterChatPayload struct {
	ChatName     string `json:"chat_name"`
	ChatMemberId string `json:"chat_member_id"`
}

type ChatStore interface {
	GetChatByName(chatName string) (*Chat, error)
	CreateChat(Chat) error
}
