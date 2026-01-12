package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"test-web/database"
	"test-web/models"

	"github.com/gorilla/mux"
)

// Главная страница: список чатов
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT ChatID, ChatName, Context FROM chats")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var chats []models.Chat
	for rows.Next() {
		var chat models.Chat
		err := rows.Scan(&chat.ChatID, &chat.ChatName, &chat.Context)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Scan error chats", http.StatusInternalServerError)
			return
		}
		chats = append(chats, chat)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, chats)
}

func ChatMessagesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatIDStr := vars["id"]
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	query := `
	SELECT m.MessageID, m.ChatID, m.UserID, m.Message, m.FileURL, m.Data, m.ReplyTo,
	       u.UserName
	FROM messages m
	JOIN users u ON m.UserID = u.UserID
	WHERE m.ChatID = $1
	ORDER BY m.Data ASC
	`

	rows, err := database.DB.Query(query, chatID)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []models.MessageWithUser
	for rows.Next() {
		var (
			messageID int
			chatID    int64
			userID    int
			message   string
			fileURL   string
			data      time.Time
			replyTo   sql.NullInt32
			userName  string
		)

		err := rows.Scan(
			&messageID,
			&chatID,
			&userID,
			&message,
			&fileURL,
			&data,
			&replyTo,
			&userName,
		)
		if err != nil {
			log.Println("Scan error:", err)
			http.Error(w, "Error reading message data", http.StatusInternalServerError)
			return
		}

		messages = append(messages, models.MessageWithUser{
			Message: models.Message{
				MessageID: messageID,
				ChatID:    chatID,
				UserID:    userID,
				Message:   message,
				FileURL:   fileURL,
				Data:      data,
				ReplyTo:   nil, // Можно доработать поддержку reply
			},
			UserName:      userName,
			FormattedTime: data.Format("02 Jan 2006, 15:04"),
		})
	}

	tmpl, err := template.ParseFiles("templates/chat.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, messages)
}
