package bot

import (
	"time"

	"github.com/ghostsama2503/alice/bot/middlewares"
	"github.com/ghostsama2503/alice/common/config"

	"github.com/aquagram/aquagram"
	"github.com/charmbracelet/log"
)

func Init() error {
	bot := aquagram.NewBot(config.Env.BotToken)

	// middlewares
	bot.Use(middlewares.BaseSettingsMiddleware)

	// handlers
	bot.OnCommand("start", StartCommandHandler)

	// init
	go logInit(bot)
	return bot.StartPolling(false)
}

func logInit(bot *aquagram.Bot) {
	for {
		if bot.Me == nil {
			time.Sleep(time.Millisecond * 100)
			continue
		}

		log.Infof("client @%s started! 🤖", bot.Me.Username)
		break
	}
}
