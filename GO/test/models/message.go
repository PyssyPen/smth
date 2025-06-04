package models

import "time"

type MessageWithUser struct {
	Message
	UserName      string
	FormattedTime string
}

type Message struct {
	MessageID int
	ChatID    int64
	UserID    int
	Message   string
	FileURL   string
	Data      time.Time
	ReplyTo   *int
}
