package database

import (
	"alice/common/config"
	"alice/database/queries"
	"database/sql"

	"github.com/charmbracelet/log"
	_ "github.com/tursodatabase/go-libsql"
)

var Db *sql.DB
var Query *queries.Queries
var connected bool

func Init() (*sql.DB, error) {

	db, err := sql.Open("libsql", config.Env.DatabaseUrl)
	if err != nil {
		return db, err
	}

	Db = db
	Query = queries.New(db)
	connected = true

	return db, nil
}

func CheckConnection() (*sql.DB, error) {

	var err error

	if !connected {
		config.LoadEnv("../.env")

		Db, err = Init()
	}

	if err != nil {
		log.Error("error checking database connection", "err", err)
	}

	return Db, err
}
