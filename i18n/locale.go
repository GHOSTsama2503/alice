package i18n

type Locale struct {
	ChatSettings           string `toml:"chat_settings"`
	CustomizeMessage       string `toml:"customize_message"`
	Disable                string `toml:"disable"`
	Enable                 string `toml:"enable"`
	OK                     string `toml:"ok"`
	SavedMessage           string `toml:"saved_message"`
	SendNewWelcomeMessage  string `toml:"send_new_welcome_message"`
	StartMessage           string `toml:"start_message"`
	TextIsRequired         string `toml:"text_is_required"`
	Welcome                string `toml:"welcome"`
	WelcomeMessageRequired string `toml:"welcome_message_required"`
	WelcomeSettings        string `toml:"welcome_settings"`
}
