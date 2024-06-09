package main

import (
	"alice/client"
	"alice/env"
	"alice/handlers"
	"alice/i18n"

	"github.com/charmbracelet/log"
	"gopkg.in/telebot.v3"
)

func main() {
	var err error

	if err = env.Load(); err != nil {
		log.Fatal(err)
	}

	if err = i18n.Load("locales"); err != nil {
		log.Fatal(err)
	}

	var bot *telebot.Bot
	if bot, err = client.Setup(); err != nil {
		log.Fatal(err)
	}

	handlers.Use(bot)

	log.Info("client started! 🤖")
	bot.Start()
}
