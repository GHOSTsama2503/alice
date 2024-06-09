package main

import (
	"alice/client"
	"alice/common/config"
	"alice/database"
	"alice/i18n"

	"github.com/charmbracelet/log"
	"gopkg.in/telebot.v3"
)

func main() {
	var err error

	if err = config.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	if err = i18n.Init("./locales"); err != nil {
		log.Fatal(err)
	}

	var bot *telebot.Bot
	if bot, err = client.Setup(); err != nil {
		log.Fatal(err)
	}

	database.Init()

	log.Infof("client @%s started! 🤖", bot.Me.Username)
	bot.Start()
}
