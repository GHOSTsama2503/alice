package client

import (
	"alice/common/config"
	"alice/handlers"
	"alice/middlewares"
	"time"

	"github.com/charmbracelet/log"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func recoverFunc(err error) {
	log.Error(err)
}

func Setup() (bot *telebot.Bot, err error) {

	botSettings := telebot.Settings{
		Token:  config.Env.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err = telebot.NewBot(botSettings)
	if err != nil {
		return
	}

	bot.Use(middleware.Recover(recoverFunc))

	middlewares.Use(bot)
	handlers.Use(bot)

	return
}
