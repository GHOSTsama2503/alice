package handlers

import (
	"alice/env"
	"alice/i18n"

	"gopkg.in/telebot.v3"
)

func StartHandler(c telebot.Context) (err error) {

	sender := c.Sender()

	var msg string
	opts := i18n.Options{
		"user": sender.FirstName,
		"me":   env.ClientName,
	}
	if msg, err = i18n.T2("start_message", sender.LanguageCode, opts); err != nil {
		return
	}

	if err = c.Reply(msg); err != nil {
		return
	}

	return
}
