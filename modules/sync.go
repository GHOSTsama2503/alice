package modules

import (
	"context"

	"github.com/aquagram/aquagram"
	"github.com/ghostsama2503/alice/database/models"
)

func SyncChat(ctx context.Context, data *aquagram.Chat) error {
	if data == nil {
		return nil
	}

	chat, err := GetChat(ctx, data.ID)
	if err != nil {
		return err
	}

	if chat.OK {
		return nil
	}

	newChat := models.Chat{
		ID:   data.ID,
		Type: data.Type,
	}

	_, err = CreateChat(ctx, newChat)

	return err
}

func SyncUser(ctx context.Context, user *aquagram.User) error {
	return nil
}
