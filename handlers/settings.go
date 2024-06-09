package handlers

import (
	"alice/i18n"
	"alice/modules"

	"github.com/charmbracelet/log"
	"gopkg.in/telebot.v3"
)

func SettingsHandler(c telebot.Context) error {

	userLocale := modules.GetUserLocale(c.Sender().ID)

	welcomeText, err := i18n.T(i18n.Welcome, userLocale)
	if err != nil {
		log.Error("error getting locale", "key", i18n.Welcome, "err", err)
		return err
	}

	firstRow := []telebot.InlineButton{
		{Text: welcomeText, Data: WelcomeSettingsCallback},
	}

	markup := &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{firstRow},
	}

	msgTextOpts := i18n.Options{
		"chat": c.Chat().Title,
	}

	msgText, err := i18n.T2(i18n.ChatSettings, userLocale, msgTextOpts)
	if err != nil {
		log.Error("error getting locale", "key", i18n.ChatSettings, "err", err)
		return err
	}

	if err := c.Reply(msgText, markup); err != nil {
		log.Error("error sending settings message", "err", err)
		return err
	}

	return nil
}
