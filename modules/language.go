package modules

import (
	"context"

	"github.com/aquagram/aquagram"
	"github.com/ghostsama2503/alice/database"
	"github.com/ghostsama2503/alice/i18n"
)

func GetLocale(ctx context.Context, event aquagram.Event) *i18n.Locale {
	chat := event.GetChat()
	from := event.GetFrom()

	if chat == nil || from == nil {
		return i18n.GetLocale(i18n.Fallback())
	}

	if chat.IsPrivate() {
		return i18n.GetLocale(GetUserLanguage(ctx, event.GetFrom()))
	}

	return i18n.GetLocale(GetChatLanguage(ctx, chat.ID))
}

const getChatLanguage = `
SELECT language_code FROM chats WHERE id = $1;
`

func GetChatLanguage(ctx context.Context, id int64) i18n.LanguageCode {
	var language string

	err := database.DB.GetContext(ctx, language, getChatLanguage, id)
	if err != nil {
		logger.Error("GetChatLanguageCode", "err", err)

		return i18n.Fallback()
	}

	return language
}

const getUserLanguage = `
SELECT language_code FROM users WHERE id = $1;
`

func GetUserLanguage(ctx context.Context, user *aquagram.User) i18n.LanguageCode {
	if user == nil {
		return i18n.Fallback()
	}

	var language string

	err := database.DB.GetContext(ctx, language, getUserLanguage, user.ID)
	if err != nil {
		logger.Error("GetUserLanguageCode", "err", err)

		return i18n.Fallback()
	}

	return language
}
