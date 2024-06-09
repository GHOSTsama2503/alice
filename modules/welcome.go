package modules

import (
	"alice/database"
	"context"

	"github.com/charmbracelet/log"
)

func IsWelcomeMessageEnabled(chatId int64) (enabled bool, err error) {

	_, err = database.CheckConnection()
	if err != nil {
		return
	}

	var result int64

	result, err = database.Query.IsWelcomeMessageEnabled(context.Background(), chatId)
	if err != nil {
		log.Error("error checking if welcome message is enabled", "chat", chatId)
		return
	}

	if result == 1 {
		enabled = true
	}

	return
}

func EnableWelcomeMessage(chatId int64) (err error) {

	_, err = database.CheckConnection()
	if err != nil {
		return
	}

	err = database.Query.EnableWelcomeMessage(context.Background(), chatId)
	if err != nil {
		log.Error("error enabling welcome message", "err", err)
	}

	return
}

func DisableWelcomeMessage(chatId int64) (err error) {

	_, err = database.CheckConnection()
	if err != nil {
		return
	}

	err = database.Query.DisableWelcomeMessage(context.Background(), chatId)
	if err != nil {
		log.Error("error disabling welcome message", "err", err)
	}

	return
}
