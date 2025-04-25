package main

import (
	"fmt"
	"log"
	"os"
	"tg_log_bot/key"
	"time"

	"gopkg.in/telebot.v4"
)

func main() {
	token, err := key.ReadKey("/home/pyssy/VSC/GO/bots/m_test_bot_2/pkg/key/key.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Замените 'YOUR_BOT_TOKEN' на токен вашего бота
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Создаем папку для логов, если она не существует
	logDir := "/home/pyssy/VSC/GO/bots/tg_log_bot/logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Обработчик для всех текстовых сообщений
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		logMessage(c.Message(), logDir)
		return nil
	})

	// Запускаем бота
	bot.Start()
}

func logMessage(m *telebot.Message, logDir string) {
	// Формируем имя файла на основе даты и времени
	fileName := fmt.Sprintf("%s/%s.txt", logDir, time.Now().Format("2006-01-02_15-04-05"))

	// Открываем файл для записи
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Failed to create log file: %v", err)
		return
	}
	defer file.Close()

	// Формируем содержимое лога
	user := m.Sender.Username
	if user == "" {
		user = fmt.Sprintf("%s %s", m.Sender.FirstName, m.Sender.LastName)
	}
	logContent := fmt.Sprintf("User: %s\nMessage: %s\n", user, m.Text)

	// Записываем лог в файл
	if _, err := file.WriteString(logContent); err != nil {
		log.Printf("Failed to write to log file: %v", err)
	}
}
