package bot

import (
	"fmt"
	"os"
	"tg_log_bot/logs"
	"time"

	"gopkg.in/telebot.v4"
)

func StartBot(token string) error {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return err
	}

	// Создаем папку для логов, если она не существует
	logDirTxt := "/home/pyssy/VSC/GO/bots/tg_log_bot/logs/txt"
	logDirJpg := "/home/pyssy/VSC/GO/bots/tg_log_bot/logs/jpg"
	if err := os.MkdirAll(logDirTxt, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(logDirJpg, 0755); err != nil {
		return err
	}

	// ВИДИМО ПРИДЕТСЯ ПИСАТЬ ХЕНДЛ ДЛЯ КАЖДОГО ДЕЙСТВИЯ

	// Обработчик для всех текстовых сообщений
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		logs.LogTxtMessage(c.Message(), logDirTxt)
		return nil
	})

	// Обработчик для всех изображений
	bot.Handle(telebot.OnPhoto, func(c telebot.Context) error {
		logs.LogJpgMessage(bot, c.Message(), logDirJpg, logDirTxt)
		return nil
	})

	// Запускаем бота
	fmt.Println("Bot has started!")
	bot.Start()
	return nil
}
