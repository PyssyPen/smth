package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/telebot.v4"
)

/*
// Вариант оптимизации, но пока вываливается в панику


	func logger(m *telebot.Message) (user, chatType, chatTitle, forwardedFrom, replyToMessage string) {
		// Формируем имя пользователя
		user = m.Sender.Username
		if user == "" {
			user = fmt.Sprintf("%s %s", m.Sender.FirstName, m.Sender.LastName)
		}

		// Определяем тип чата
		chatType = "Unknown"
		chatTitle = ""
		if m.Chat != nil {
			switch m.Chat.Type {
			case telebot.ChatPrivate:
				chatType = "Private"
			case telebot.ChatGroup:
				chatType = "Group"
			case telebot.ChatSuperGroup:
				chatType = "Supergroup"
			case telebot.ChatChannel:
				chatType = "Channel"
			}
			chatTitle = m.Chat.Title
		}

		// Формируем информацию о пересланном сообщении
		if m.Sender.Username != "" {
			forwardedFrom = fmt.Sprintf("Sender Name: %s", m.Sender.Username)
		} else if m.SenderChat != nil {
			forwardedFrom = fmt.Sprintf("Chat ID: %d, Title: %s", m.SenderChat.ID, m.SenderChat.Title)
		}

		// Формируем информацию о сообщении, на которое ответили
		if m.ReplyTo != nil {
			replyToMessage = fmt.Sprintf("Message ID: %d, Text: %s", m.ReplyTo.ID, m.ReplyTo.Text)
		} else {
			replyToMessage = " - "
		}

		return user, chatType, chatTitle, forwardedFrom, replyToMessage
	}


	// Использование:
	user, chatType, chatTitle, forwardedFrom, replyToMessage = logger(&telebot.Message{})
*/

func LogTxtMessage(m *telebot.Message, logDir string) {

	text := m.Text
	txtName := fmt.Sprintf("%s/%s.txt", logDir, time.Now().Format("2006-01-02_15-04-05"))

	// Открываем файл для записи
	txtfile, err := os.Create(txtName)
	if err != nil {
		log.Printf("Failed to create log file: %v", err)
		return
	}
	defer txtfile.Close()

	// Формируем имя пользователя
	user := m.Sender.Username
	if user == "" {
		user = fmt.Sprintf("%s %s", m.Sender.FirstName, m.Sender.LastName)
	}

	// Определяем тип чата
	chatType := "Unknown"
	chatTitle := ""
	if m.Chat != nil {
		switch m.Chat.Type {
		case telebot.ChatPrivate:
			chatType = "Private"
		case telebot.ChatGroup:
			chatType = "Group"
		case telebot.ChatSuperGroup:
			chatType = "Supergroup"
		case telebot.ChatChannel:
			chatType = "Channel"
		}
		chatTitle = m.Chat.Title
	}

	// Формируем информацию о пересланном сообщении
	var forwardedFrom string
	if m.Sender.Username != "" {
		forwardedFrom = fmt.Sprintf("Sender Name: %s", m.Sender.Username)
	} else if m.SenderChat != nil {
		forwardedFrom = fmt.Sprintf("Chat ID: %d, Title: %s", m.SenderChat.ID, m.SenderChat.Title)
	}

	// Формируем информацию о сообщении, на которое ответили
	var replyToMessage string
	if m.ReplyTo != nil {
		replyToMessage = fmt.Sprintf("Message ID: %d, Text: %s", m.ReplyTo.ID, m.ReplyTo.Text)
	} else {
		replyToMessage = " - "
	}

	photoFileID := "-"
	imagePath := "-"
	// Формируем содержимое лога
	logContent := fmt.Sprintf(
		"Message ID: %d\nChat ID: %d\nChat Type: %s\nChat Title: %s\nUser: %s\nUser ID: %d\nMessage: %s\nPhoto File ID: %s\nPhoto Path: %s\nSent At: %s\nForwarded From: %s\nReply To Message: %s\n",
		m.ID,
		m.Chat.ID,
		chatType,
		chatTitle,
		user,
		m.Sender.ID,
		text,
		photoFileID,
		imagePath,
		m.Time().Format(time.RFC3339),
		forwardedFrom,
		replyToMessage,
	)

	// Записываем лог в файл
	if _, err := txtfile.WriteString(logContent); err != nil {
		log.Printf("Failed to write to log file: %v", err)
	}
}

func LogJpgMessage(bot *telebot.Bot, m *telebot.Message, logDirJpg string, logDirTxt string) {

	Text := m.Caption
	if Text == "" {
		Text = "Нет подписи"
	}

	// Получаем информацию о фото
	photo := m.Photo

	file, err := bot.File(&telebot.File{FileID: photo.FileID})
	if err != nil {
		log.Printf("Failed to get file info: %v", err)
		return
	}

	// Формируем имя файла для изображения
	imageName := fmt.Sprintf("IMG_%s.jpg", time.Now().Format("2006-01-02_15-04-05"))
	imagePath := filepath.Join(logDirJpg, imageName)
	f, err := os.Create(imagePath)
	if err != nil {
		log.Printf("Failed to create image file: %v", err)
		return
	}
	defer f.Close()

	// Копируем содержимое файла в созданный файл
	_, err = io.Copy(f, file)
	if err != nil {
		log.Printf("Failed to copy file content: %v", err)
		return
	}

	// Формируем имя файла для лога на основе текущей даты и времени
	logName := fmt.Sprintf("%s/%s.txt", logDirTxt, time.Now().Format("2006-01-02_15-04-05"))
	logFile, err := os.Create(logName)
	if err != nil {
		log.Printf("Failed to create log file: %v", err)
		return
	}
	defer logFile.Close()

	// Формируем имя пользователя
	user := m.Sender.Username
	if user == "" {
		user = fmt.Sprintf("%s %s", m.Sender.FirstName, m.Sender.LastName)
	}

	// Определяем тип чата
	chatType := "Unknown"
	chatTitle := ""
	if m.Chat != nil {
		switch m.Chat.Type {
		case telebot.ChatPrivate:
			chatType = "Private"
		case telebot.ChatGroup:
			chatType = "Group"
		case telebot.ChatSuperGroup:
			chatType = "Supergroup"
		case telebot.ChatChannel:
			chatType = "Channel"
		}
		chatTitle = m.Chat.Title
	}

	// Формируем информацию о пересланном сообщении
	var forwardedFrom string
	if m.Sender.Username != "" {
		forwardedFrom = fmt.Sprintf("Sender Name: %s", m.Sender.Username)
	} else if m.SenderChat != nil {
		forwardedFrom = fmt.Sprintf("Chat ID: %d, Title: %s", m.SenderChat.ID, m.SenderChat.Title)
	}

	// Формируем информацию о сообщении, на которое ответили
	var replyToMessage string
	if m.ReplyTo != nil {
		replyToMessage = fmt.Sprintf("Message ID: %d, Text: %s", m.ReplyTo.ID, m.ReplyTo.Text)
	} else {
		replyToMessage = " - "
	}

	// Формируем содержимое лога
	logContent := fmt.Sprintf(
		"Message ID: %d\nChat ID: %d\nChat Type: %s\nChat Title: %s\nUser: %s\nUser ID: %d\nMessage: %s\nPhoto File ID: %s\nPhoto Path: %s\nSent At: %s\nForwarded From: %s\nReply To Message: %s\n",
		m.ID,
		m.Chat.ID,
		chatType,
		chatTitle,
		user,
		m.Sender.ID,
		Text,
		photo.FileID,
		imagePath,
		m.Time().Format(time.RFC3339),
		forwardedFrom,
		replyToMessage,
	)

	// Записываем лог в файл
	if _, err := logFile.WriteString(logContent); err != nil {
		log.Printf("Failed to write to log file: %v", err)
	}

	fmt.Println("JPG logged and saved")
}
