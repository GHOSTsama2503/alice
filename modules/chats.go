package modules

import (
	"context"
	"database/sql"

	"github.com/ghostsama2503/alice/database"
	"github.com/ghostsama2503/alice/database/models"

	"github.com/charmbracelet/log"
)

const createChat = `
INSERT INTO chats (id, type, title)
VALUES ($1, $2, $3)
RETURNING id, type, title, created_at, updated_at;
`

const createChatSettings = `
INSERT INTO chat_settings (id)
VALUES ($1);
`

func CreateChat(ctx context.Context, data models.Chat) (models.Chat, error) {
	tx, err := database.DB.BeginTxx(ctx, nil)
	if err != nil {
		return data, err
	}

	defer tx.Rollback()

	row := tx.QueryRowxContext(ctx, createChat, data.ID, data.Type, data.Title)

	var chat models.Chat

	err = row.StructScan(&chat)
	if err != nil {
		return chat, err
	}

	_, err = tx.ExecContext(ctx, createChatSettings, chat.ID)
	if err != nil {
		return chat, err
	}

	return chat, tx.Commit()
}

const getChat = `
SELECT * FROM chats WHERE id = $1;
`

func GetChat(ctx context.Context, id int64) (Optional[models.Chat], error) {
	var chat Optional[models.Chat]

	err := database.DB.GetContext(ctx, &chat.Value, getChat, id)
	if err == sql.ErrNoRows {
		return chat, nil
	}

	if err != nil {
		log.Error("GetChat", "id", id, "err", err)
		return chat, err
	}

	chat.OK = true
	return chat, nil
}

const getChatSettings = `
SELECT * FROM chat_settings WHERE chat_id = $1;
`

func GetChatSettings(ctx context.Context, id int64) (Optional[models.ChatSettings], error) {
	var chatSettings Optional[models.ChatSettings]

	err := database.DB.GetContext(ctx, &chatSettings.Value, getChatSettings, id)
	if err == sql.ErrNoRows {
		return chatSettings, nil
	}

	if err != nil {
		log.Error("GetChatSettings", "id", id, "err", err)
		return chatSettings, err
	}

	return chatSettings, nil
}
