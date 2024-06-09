package handlers

import (
	"alice/env"
	"alice/i18n"
	"strings"

	"gopkg.in/telebot.v3"
)

func StartHandler(c telebot.Context) (err error) {

	sender := c.Sender()

	locale := i18n.GetLocale(sender.LanguageCode)

	replacer := strings.NewReplacer("{user}", sender.FirstName, "{me}", env.ClientName)
	msg := replacer.Replace(locale.StartMessagee)

	if err = c.Reply(msg); err != nil {
		return
	}

	return
}
