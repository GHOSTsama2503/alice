package handlers

import "gopkg.in/telebot.v3"

var Bot *telebot.Bot

func Use(bot *telebot.Bot) {
	Bot = bot

	bot.Handle("/start", StartHandler)
}
