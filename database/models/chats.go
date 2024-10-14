package models

import (
	"time"

	"github.com/aquagram/aquagram"
)

type Chat struct {
	ID        int64             `db:"id"`
	Type      aquagram.ChatType `db:"type"`
	Title     string            `db:"title"`
	CreatedAt time.Time         `db:"created_at"`
	UpdatedAt time.Time         `db:"updated_at"`
}

type ChatSettings struct {
	ID                      int64 `db:"id"`
	IsCaptchaEnabled        bool  `db:"is_captcha_enabled"`
	IsWelcomeMessageEnabled bool  `db:"is_welcome_message_enabled"`
}
