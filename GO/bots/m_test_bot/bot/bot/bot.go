package bot

import (
	"fmt"
	"time"

	"m_test_bot/bot/handlers"

	tele "gopkg.in/telebot.v4"
)

func StartBot(token string) error {
	settings := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(settings)
	if err != nil {
		return err
	}

	handlers.RegisterHandlers(bot)

	fmt.Println("Бот запущен...")
	bot.Start()
	return nil
}
