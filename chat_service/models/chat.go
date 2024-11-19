package models

import "time"

type Chat struct {
	ChatID   int    `json:"chat_id"`
	ChatName string `json:"chat_name"`
}

type ChatMember struct {
	ChatMemberID int       `json:"chat_member_id"`
	ChatID       int       `json:"chat_id"`
	UserID       int       `json:"user_id"`
	JoinedAt     time.Time `json:"joined_at"`
}

type ChatMessage struct {
	MessageID   int       `json:"message_id"`
	ChatID      int       `json:"chat_id"`
	UserID      int       `json:"user_id"`
	MessageText string    `json:"message_text"`
	SentAt      time.Time `json:"sent_at"`
}

type RegisterChatPayload struct {
	ChatName     string `json:"chat_name"`
	ChatMemberId string `json:"chat_member_id"`
}

type ChatStore interface {
	GetChatByName(chatName string) (*Chat, error)
	CreateChat(Chat) error
	GetUserChats(userID int) ([]Chat, error)
}
