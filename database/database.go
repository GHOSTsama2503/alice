package database

import (
	"github.com/ghostsama2503/alice/common/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init() (*sqlx.DB, error) {
	var err error

	DB, err = sqlx.Open("pgx", config.Env.DatabaseURL)
	if err != nil {
		return DB, err
	}

	return DB, nil
}
