package bot

import (
	"github.com/ghostsama2503/alice/common/config"
	"github.com/ghostsama2503/alice/i18n"

	"github.com/aquagram/aquagram"
)

func StartCommandHandler(_ *aquagram.Bot, message *aquagram.Message) error {
	locale, err := i18n.GetLocale(message.From.LanguageCode)
	if err != nil {
		return err
	}

	opts := i18n.Options{
		"user": message.From.TextMention(aquagram.ParseModeHtml),
		"me":   config.Env.ClientName,
	}

	text := i18n.WithOptions(locale.StartMessage, opts)

	_, err = message.Reply(text, &aquagram.SendMessageParams{
		ParseMode: aquagram.ParseModeHtml,
	})

	return err
}
