package env

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	ClientName string

	BotToken      string
	IsDevelopment bool
)

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	ClientName = os.Getenv("CLIENT_NAME")
	if ClientName == "" {
		return errors.New("empty client name")
	}

	BotToken = os.Getenv("BOT_TOKEN")
	if BotToken == "" {
		return errors.New("empty bot token")
	}

	environment := strings.ToLower(os.Getenv("ENVIRONMENT"))
	if environment == "development" || environment == "develop" || environment == "devel" || environment == "dev" {
		IsDevelopment = true
	}

	return nil
}
