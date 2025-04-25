package handlers

import (
	"fmt"
	keyboard "m_test_bot_2/pkg/keyboards"

	tele "gopkg.in/telebot.v4"
)

func RegisterCallback(bot *tele.Bot) {
	bot.Handle(&keyboard.Steak, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Вырезка"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали вырезку!", nil)
		ZeroString = "Вырезка"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Ribs, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Ребра"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали ребра!", nil)
		ZeroString = "Ребра"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Mince, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Фарш"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали фарш!", nil)
		ZeroString = "Фарш"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Ribeye, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Рибай"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали рибай!", nil)
		ZeroString = "Рибай"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Liver, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Печень"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали печень!", nil)
		ZeroString = "Печень"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.FrontThigh, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Ляжка передняя"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали ляжку переднюю!", nil)
		ZeroString = "Ляжка передняя"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.RearThigh, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Ляжка задняя"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали ляжку заднюю!", nil)
		ZeroString = "Ляжка задняя"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Lamb, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Баранина"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали баранину!", nil)
		ZeroString = "Баранина"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Goose, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Гусь"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали гуся!", nil)
		ZeroString = "Гусь"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Duck, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Утка"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали утку!", nil)
		ZeroString = "Утка"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.HorseMeat, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Конина"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали конину!", nil)
		ZeroString = "Конина"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Chicken, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Курица"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали курицу!", nil)
		ZeroString = "Курица"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Kazylyk, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Казылык"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали казылык!", nil)
		ZeroString = "Казылык"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Turkey, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Тур"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали тур!", nil)
		ZeroString = "Тур"
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Tutyrma, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Тутырма"); err != nil {
			return fmt.Errorf("ошибка логирования: %v", err)
		}
		sendMessage(c, "Вы выбрали тутырму!", nil)
		ZeroString = "Тутырма"
		return KgHandler(ZeroString, c)
	})
}
