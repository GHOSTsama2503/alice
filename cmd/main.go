package main

import (
	"alice/client"
	"alice/common/config"
	"alice/database"
	"alice/i18n"
	"alice/modules"

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

	_, err = database.Init()
	if err != nil {
		log.Fatal(err)
	}

	if err = modules.Init(); err != nil {
		log.Fatal(err)
	}

	log.Infof("client @%s started! 🤖", bot.Me.Username)
	bot.Start()
}
