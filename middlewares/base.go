package middlewares

import (
	"alice/common/logerr"
	"alice/database"
	"alice/database/queries"
	"context"

	"gopkg.in/telebot.v3"
)

func BaseSettingsMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) (err error) {

		sender := c.Sender()

		_, err = database.Query.GetUserLocale(context.Background(), sender.ID)
		if err == nil {
			return next(c)
		}

		params := queries.CreateUserParams{
			ID:           sender.ID,
			LanguageCode: sender.LanguageCode,
		}

		_, err = database.Query.CreateUser(context.Background(), params)
		if err != nil {
			return logerr.Error(err)
		}

		return next(c)
	}
}
