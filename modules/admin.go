package modules

import (
	"gopkg.in/telebot.v3"
)

func IsChatAdmin(c telebot.Context) (bool, error) {

	bot := c.Bot()
	chat := c.Chat()
	userId := c.Sender().ID

	var isAdmin bool

	admins, err := bot.AdminsOf(chat)
	if err != nil {
		return isAdmin, err
	}

	for _, admin := range admins {
		if userId == admin.User.ID {
			isAdmin = true
			break
		}
	}

	return isAdmin, nil
}
