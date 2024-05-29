package modules

import (
	"alice/common/config"
	"alice/database"
	"alice/database/queries"
	"context"
	"fmt"

	"github.com/charmbracelet/log"
)

func GetUserLocale(uid int64) string {

	database.CheckConnection()

	locale, err := database.Query.GetUserLocale(context.Background(), uid)
	if err != nil {
		errMsg := fmt.Sprintf("error getting user locale, using fallback: '%s'", config.Env.DefaultLocale)
		log.Warn(errMsg, "err", err)

		return config.Env.DefaultLocale
	}
	return locale
}

func SetUserLocale(uid int64, locale string) error {

	params := queries.SetUserLocaleParams{
		ID:           uid,
		LanguageCode: locale,
	}

	return database.Query.SetUserLocale(context.Background(), params)
}
