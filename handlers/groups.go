package handlers

import (
	"alice/common/logerr"
	"alice/database"
	"alice/database/queries"
	"context"

	"github.com/charmbracelet/log"
	"gopkg.in/telebot.v3"
)

func AddedToGroupHandler(c telebot.Context) error {

	chat := c.Chat()
	log.Info("added to group", "id", chat.ID, "title", chat.Title)

	var err error

	_, err = database.CheckConnection()
	if err != nil {
		return logerr.Error(err)
	}

	result, err := database.Query.GroupExists(context.Background(), chat.ID)
	if err != nil {
		log.Error("error checking if group is already on db", "id", chat.ID, "err", err)
		return err
	}

	if len(result) == 1 {
		return nil
	}

	tx, err := database.Db.Begin()
	if err != nil {
		return err
	}

	query := database.Query.WithTx(tx)
	ctx := context.Background()

	var isPrivate int64
	if chat.Private {
		isPrivate = 1
	}

	groupParams := queries.CreateGroupParams{
		ID:        chat.ID,
		Title:     chat.Title,
		IsPrivate: isPrivate,
	}

	err = query.CreateGroup(ctx, groupParams)
	if err != nil {
		_ = tx.Rollback()
		return logerr.Error(err)
	}

	settingsParams := queries.CreateGroupSettingsParams{
		GroupID:                 chat.ID,
		IsWelcomeMessageEnabled: 0,
	}

	err = query.CreateGroupSettings(ctx, settingsParams)
	if err != nil {
		_ = tx.Rollback()
		return logerr.Error(err)
	}

	messagesParams := queries.CreateMessagesParams{
		GroupID: chat.ID,
	}

	err = query.CreateMessages(ctx, messagesParams)
	if err != nil {
		_ = tx.Rollback()
		return logerr.Error(err)
	}

	if err = tx.Commit(); err != nil {
		return logerr.Error(err)
	}

	return nil
}
