package middlewares

import (
	"github.com/aquagram/aquagram"
	"github.com/ghostsama2503/alice/modules"
)

func BaseSettingsMiddleware(bot *aquagram.Bot, event aquagram.Event) error {
	err := modules.SyncChat(bot.Context(), event.GetChat())
	if err != nil {
		return err
	}

	err = modules.SyncUser(bot.Context(), event.GetFrom())
	if err != nil {
		return err
	}

	return nil
}
