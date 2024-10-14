package models

import (
	"time"

	"github.com/ghostsama2503/alice/i18n"
)

type User struct {
	ID           int64             `db:"id"`
	LanguageCode i18n.LanguageCode `db:"language_code"`
	CreatedAt    time.Time         `db:"created_at"`
	UpdatedAt    time.Time         `db:"updated_at"`
}

type UsersStats struct {
	ID       int64 `db:"id"`
	ChatID   int64 `db:"chat_id"`
	UserID   int64 `db:"user_id"`
	Warnings int   `db:"warnings"`
}
