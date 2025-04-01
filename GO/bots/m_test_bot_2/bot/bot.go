package bot

import (
	"fmt"
	"time"

	"m_test_bot_2/pkg/handlers"

	tele "gopkg.in/telebot.v4"
)

func StartBot(token string) error {
	settings := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	Bot, err := tele.NewBot(settings)
	if err != nil {
		return err
	}

	handlers.RegisterHandlers(Bot)
	handlers.RegisterCallback(Bot)

	fmt.Println("Бот запущен...")
	Bot.Start()
	return nil
}
