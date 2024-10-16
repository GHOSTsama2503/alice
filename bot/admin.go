package bot

import "github.com/aquagram/aquagram"

func WarningMessageHandler(bot *aquagram.Bot, message *aquagram.Message) error {
	return nil
}

func KickCommandHandler(bot *aquagram.Bot, message *aquagram.Message) error {
	// TODO kick feature

	return nil
}

func BanCommandHandler(bot *aquagram.Bot, message *aquagram.Message) error {
	// TODO ban and send a message that contains an "unban" button

	return nil
}

func UnbanCommandHandler(bot *aquagram.Bot, message *aquagram.Message) error {
	// TODO unban feature

	return nil
}

func UnbanCallbackHandler(bot *aquagram.Bot, callback *aquagram.CallbackQuery) error {
	// TODO unban feature

	return nil
}
