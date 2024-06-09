package modules

import (
	"alice/database"
	"alice/database/queries"
	"context"

	"github.com/charmbracelet/log"
)

func GetWelcomeMessage(chatId int64) (message string, err error) {

	_, err = database.CheckConnection()
	if err != nil {
		return
	}

	message, err = database.Query.GetWelcomeMessage(context.Background(), chatId)
	if err != nil {
		log.Error("error getting welcome message", "err", err)
		return
	}

	return
}

func SetWelcomeMessage(chatId int64, message string) (err error) {

	_, err = database.CheckConnection()
	if err != nil {
		return
	}

	params := queries.SetWelcomeMessageParams{
		GroupID:        chatId,
		WelcomeMessage: message,
	}

	err = database.Query.SetWelcomeMessage(context.Background(), params)
	if err != nil {
		log.Error("error setting welcome message", "err", err)
	}

	return
}
