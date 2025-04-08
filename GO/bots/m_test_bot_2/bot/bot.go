package bot

import (
	"fmt"
	"m_test_bot_2/pkg/handlers"
	"time"

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
	Bot.Use(func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			go func() {
				if err := next(c); err != nil {
					fmt.Printf("Ошибка обработки: %v, UserID: %d\n", err, c.Sender().ID)
				}
			}()
			return nil
		}
	})
	handlers.RegisterHandlers(Bot)
	handlers.RegisterCallback(Bot)
	fmt.Println("Бот запущен...")
	Bot.Start()
	return nil
}
