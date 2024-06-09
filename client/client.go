package client

import (
	"alice/env"
	"time"

	"gopkg.in/telebot.v3"
)

func Setup() (bot *telebot.Bot, err error) {

	botSettings := telebot.Settings{
		Token:  env.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err = telebot.NewBot(botSettings)
	if err != nil {
		return
	}

	return
}
