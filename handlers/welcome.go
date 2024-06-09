package handlers

import (
	"alice/common/logerr"
	"alice/i18n"
	"alice/modules"
	"fmt"

	"github.com/charmbracelet/log"
	"gopkg.in/telebot.v3"
)

func WelcomeHandler(c telebot.Context) error {

	var isEnabled bool

	chatId := c.Chat().ID

	isEnabled, err := modules.IsWelcomeMessageEnabled(chatId)
	if err != nil {
		return err
	}

	if !isEnabled {
		return nil
	}

	welcomeMessage, err := modules.GetWelcomeMessage(chatId)
	if err != nil {
		return err
	}

	if err := c.Reply(welcomeMessage); err != nil {
		log.Error("error sending welcome message", "chat", chatId, "err", err)
		return err
	}

	return nil
}

func WelcomeSettings(c telebot.Context) error {

	userLocale := modules.GetUserLocale(c.Sender().ID)

	isEnabled, err := modules.IsWelcomeMessageEnabled(c.Chat().ID)
	if err != nil {
		return err
	}

	var switchText string
	var switchCallback string

	if isEnabled {
		switchCallback = DisableWelcomeCallback

		switchText, err = i18n.T(i18n.Disable, userLocale)
		if err != nil {
			return logerr.Error(err)
		}

	} else {
		switchCallback = EnableWelcomeCallback

		switchText, err = i18n.T(i18n.Enable, userLocale)
		if err != nil {
			return logerr.Error(err)
		}
	}

	switchBtn := []telebot.InlineButton{
		{Text: modules.NewSwitchButton(!isEnabled, switchText), Data: switchCallback},
	}

	customizeText, err := i18n.T(i18n.CustomizeMessage, userLocale)
	if err != nil {
		return logerr.Error(err)
	}

	customizeBtn := []telebot.InlineButton{
		{Text: fmt.Sprintf("📝 %s", customizeText), Data: CustomizeWelcomeCallback},
	}

	markup := &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{switchBtn, customizeBtn},
	}

	msgText, err := i18n.T(i18n.WelcomeSettings, userLocale)
	if err != nil {
		log.Error("error getting locale", "key", i18n.WelcomeSettings, "err", err)
		return err
	}

	if err := c.Edit(msgText, markup); err != nil {
		log.Error("error editing welcome settings message", "err", err)
		return err
	}

	return nil
}

func WelcomeSettingsEnable(c telebot.Context) error {

	chatId := c.Chat().ID
	userId := c.Sender().ID

	userLocale := modules.GetUserLocale(userId)

	isAdmin, err := modules.IsChatAdmin(c)
	if err != nil {
		return err
	}

	if !isAdmin {
		resp := &telebot.CallbackResponse{
			Text:      "you are not admin",
			ShowAlert: true,
		}

		if err := c.Respond(resp); err != nil {
			log.Error("error responding to user callback", "err", err)
		}

		return nil
	}

	message, err := modules.GetWelcomeMessage(chatId)
	if err != nil {
		return err
	}

	if message == "" {
		errMsg, err := i18n.T(i18n.WelcomeMessageRequired, userLocale)
		if err != nil {
			log.Error("error getting locale", "key", i18n.WelcomeMessageRequired, "user", userId, "err", err)
			return err
		}

		resp := &telebot.CallbackResponse{
			Text:      errMsg,
			ShowAlert: true,
		}

		if err := c.Respond(resp); err != nil {
			log.Error("error responding to callback", "err", err)
			return err
		}

		return nil
	}

	if err := modules.EnableWelcomeMessage(chatId); err != nil {
		return err
	}

	return WelcomeSettings(c)
}

func WelcomeSettingsDisable(c telebot.Context) error {

	chatId := c.Chat().ID
	//userId := c.Sender().ID

	//userLocale := modules.GetUserLocale(userId)

	isAdmin, err := modules.IsChatAdmin(c)
	if err != nil {
		return err
	}

	if !isAdmin {
		resp := &telebot.CallbackResponse{
			Text:      "you are not admin",
			ShowAlert: true,
		}

		if err := c.Respond(resp); err != nil {
			log.Error("error responding to user callback", "err", err)
		}

		return nil
	}

	if err := modules.DisableWelcomeMessage(chatId); err != nil {
		return err
	}

	return WelcomeSettings(c)
}

func WelcomeSettingsCustomize(c telebot.Context) error {

	userLocale := modules.GetUserLocale(c.Sender().ID)

	isAdmin, err := modules.IsChatAdmin(c)
	if err != nil {
		return err
	}

	if !isAdmin {
		resp := &telebot.CallbackResponse{
			Text:      "you are not admin",
			ShowAlert: true,
		}

		if err := c.Respond(resp); err != nil {
			log.Error("error responding to user callback", "err", err)
		}

		return nil
	}

	newText, err := i18n.T(i18n.SendNewWelcomeMessage, userLocale)
	if err != nil {
		log.Error("error getting locale", "key", i18n.SendNewWelcomeMessage, "err", err)
		return err
	}

	if err := c.Edit(newText); err != nil {
		log.Error("error editing message", "err", err)
		return err
	}

	modules.SetState(c.Sender().ID, WaitingWelcomeMessage)

	return nil
}

func WaitingWelcomeMessage(c telebot.Context) error {

	userLocale := modules.GetUserLocale(c.Sender().ID)

	message := c.Text()
	if message == "" {

		errMsg, err := i18n.T(i18n.TextIsRequired, userLocale)
		if err != nil {
			log.Error("error getting locale", "key", i18n.TextIsRequired, "err", err)
			return err
		}

		if err := c.Reply(errMsg); err != nil {
			log.Error("error sending error message", "err", err)
			return err
		}

		return nil
	}

	if err := modules.SetWelcomeMessage(c.Chat().ID, message); err != nil {
		modules.DelState(c.Sender().ID)
		return err
	}

	modules.DelState(c.Sender().ID)

	savedMsg, err := i18n.T(i18n.SavedMessage, userLocale)
	if err != nil {
		log.Error("error getting locale", "key", i18n.SavedMessage, "err", err)
	}

	if err := c.Reply(savedMsg); err != nil {
		log.Error("error sending message", "err", err)
		return err
	}

	return nil
}
