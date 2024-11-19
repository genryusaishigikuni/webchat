package models

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}

}

func (c *Store) GetChatByName(chatName string) (*Chat, error) {

	rows, err := c.db.Query("SELECT chat_id, chat_name, chat_members FROM chats WHERE chat_name = ?", chatName)
	if err != nil {
		return nil, err
	}

	ch := new(Chat)
	for rows.Next() {
		ch, err = scanRowIntoChat(rows)
		if err != nil {
			return nil, err
		}
	}
	if ch.ChatID == 0 {
		return nil, fmt.Errorf("chat not found")
	}
	return ch, nil
}

func scanRowIntoChat(row *sql.Rows) (*Chat, error) {
	chat := new(Chat)
	err := row.Scan(
		&chat.ChatID,
		&chat.ChatName,
		&chat.ChatMembers)
	if err != nil {
		return nil, err
	}

	return chat, err
}

func (c *Store) CreateChat(Chat) error {
	return nil
}
