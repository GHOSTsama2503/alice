package handlers

import (
	"gopkg.in/telebot.v3"
)

var Bot *telebot.Bot

func Use(bot *telebot.Bot) {
	Bot = bot

	// a hack for enable middlewares
	bot.Handle(telebot.OnText, GlobalHandler)
	bot.Handle(telebot.OnCallback, GlobalHandler)
	bot.Handle(telebot.OnMedia, GlobalHandler)

	// events
	bot.Handle(telebot.OnAddedToGroup, AddedToGroupHandler)
	bot.Handle(telebot.OnUserJoined, WelcomeHandler)

	// commands
	bot.Handle("/start", StartHandler)
	bot.Handle("/settings", SettingsHandler)

	// callbacks
	bot.Handle(telebot.OnCallback, CallbacksHandler)
}

func CallbacksHandler(c telebot.Context) error {

	data := c.Data()

	switch data {
	case WelcomeSettingsCallback:
		WelcomeSettings(c)
	case EnableWelcomeCallback:
		WelcomeSettingsEnable(c)
	case DisableWelcomeCallback:
		WelcomeSettingsDisable(c)
	case CustomizeWelcomeCallback:
		WelcomeSettingsCustomize(c)
	}

	return nil
}
