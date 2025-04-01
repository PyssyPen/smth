package main

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token:  "",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Настраиваем основную клавиатуру и обработчики
	setupMainKeyboard(b)

	b.Start()
}

// Функция для настройки основной клавиатуры и обработчиков
func setupMainKeyboard(b *tele.Bot) {
	// Создаем основную клавиатуру
	mainKeyboard := &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	// Добавляем кнопки
	btnHelp := mainKeyboard.Text("ℹ Help")
	btnSettings := mainKeyboard.Text("⚙ Settings")

	// Размещаем кнопки в рядах
	mainKeyboard.Reply(
		mainKeyboard.Row(btnHelp),
		mainKeyboard.Row(btnSettings),
	)

	// Обработчик команды /start
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Привет! Выберите опцию:", mainKeyboard)
	})

	// Обработчик нажатия на кнопку "Help"
	b.Handle(&btnHelp, func(c tele.Context) error {
		return c.Send("Здесь можно получить помощь: ...")
	})

	// Обработчик нажатия на кнопку "Settings"
	b.Handle(&btnSettings, func(c tele.Context) error {
		// Настраиваем клавиатуру настроек
		return setupSettingsKeyboard(c, b)
	})
}

// Функция для настройки клавиатуры настроек
func setupSettingsKeyboard(c tele.Context, b *tele.Bot) error {
	mainKeyboard := &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	// Добавляем кнопки
	btnHelp := mainKeyboard.Text("ℹ Help")
	btnSettings := mainKeyboard.Text("⚙ Settings")

	// Размещаем кнопки в рядах
	mainKeyboard.Reply(
		mainKeyboard.Row(btnHelp),
		mainKeyboard.Row(btnSettings),
	)

	// Создаем клавиатуру настроек
	settingsKeyboard := &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	// Добавляем кнопки настроек
	btnProfile := settingsKeyboard.Text("👤 Profile")
	btnNotifications := settingsKeyboard.Text("🔔 Notifications")
	btnBack := settingsKeyboard.Text("⬅ Back")

	// Размещаем кнопки в рядах
	settingsKeyboard.Reply(
		settingsKeyboard.Row(btnProfile),
		settingsKeyboard.Row(btnNotifications),
		settingsKeyboard.Row(btnBack),
	)

	// Отправляем сообщение с новой клавиатурой
	if err := c.Send("Выберите настройку:", settingsKeyboard); err != nil {
		return err
	}

	// Обработчики для кнопок настроек
	b.Handle(&btnProfile, func(c tele.Context) error {
		return c.Send("Здесь можно настроить профиль: ...")
	})

	b.Handle(&btnNotifications, func(c tele.Context) error {
		return c.Send("Здесь можно настроить уведомления: ...")
	})

	b.Handle(&btnBack, func(c tele.Context) error {
		// Возвращаемся к основной клавиатуре
		setupMainKeyboard(b)
		return c.Send("", mainKeyboard) // Возврат к главному меню:
	})

	return nil
}
