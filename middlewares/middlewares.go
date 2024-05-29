package middlewares

import "gopkg.in/telebot.v3"

func Use(bot *telebot.Bot) {
	bot.Use(BaseSettingsMiddleware)
}
