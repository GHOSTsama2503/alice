package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Environment struct {
	BotToken string `env:"BOT_TOKEN"`

	// database
	DatabaseUrl    string `env:"DATABASE_URL" envDefault:"database.db"`
	DatabaseSchema string `env:"DATABASE_SCHEMA"`

	// app
	ClientName    string `env:"CLIENT_NAME"`
	DefaultLocale string `env:"DEFAULT_LOCALE" envDefault:"en"`

	IsProduction bool `env:"PRODUCTION"`
}

var Env *Environment

func LoadEnv(filenames ...string) error {
	godotenv.Load(filenames...)

	environment := &Environment{}
	if err := env.Parse(environment); err != nil {
		return err
	}

	Env = environment

	return nil
}
