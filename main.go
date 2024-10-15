package main

import (
	"github.com/ghostsama2503/alice/bot"
	"github.com/ghostsama2503/alice/common/config"
	"github.com/ghostsama2503/alice/database"
	"github.com/ghostsama2503/alice/i18n"

	"github.com/charmbracelet/log"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	if _, err := database.Init(); err != nil {
		log.Fatal(err)
	}

	if err := i18n.Init(config.Env.DefaultLocale); err != nil {
		log.Fatal(err)
	}

	if err := bot.Init(); err != nil {
		log.Fatal(err)
	}
}
