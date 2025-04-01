package handlers

import (

	//"log"
	"io/ioutil"
	keyboard "m_test_bot_2/pkg/keyboards"
	"os"
	"path/filepath"

	"fmt"

	tele "gopkg.in/telebot.v4"
)

func handleInfoCommand(c tele.Context) error {
	user := c.Sender()
	fmt.Printf("User ID: %d\n", user.ID)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("First Name: %s\n", user.FirstName)
	fmt.Printf("Last Name: %s\n", user.LastName)
	fmt.Printf("Language Code: %s\n", user.LanguageCode)
	return nil
}

func RegisterHandlers(bot *tele.Bot) {
	bot.Handle("/start", startHandler)
	bot.Handle(tele.OnText, messageHandler)
}

func startHandler(c tele.Context) error {
	fmt.Println("Получена команда /start")
	return c.Send(handleCommandStart(c, "start.txt", keyboard.CreateMenu()))
}

func handleCommandStart(c tele.Context, filePath string, keyboard *tele.ReplyMarkup) error {
	// Читаем содержимое файла
	message, err := readDescriptionFromFile(filePath)
	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка при чтении файла: %v", err))
	}
	return c.Send(message, keyboard)
}

// Функция для отправки сообщения с клавиатурой (инлайн)
func handleCommand(c tele.Context, meatName, imagePath, descriptionFilePath string, keyboard *tele.ReplyMarkup) error {
	// Отправляем название мяса без клавиатуры
	if err := c.Send(meatName); err != nil {
		return err
	}
	// Отправляем изображение
	photo := &tele.Photo{File: tele.FromDisk(filepath.Join("/home/pyssy/VSC/GO/bots/m_test_bot_2/pkg/jpg", imagePath))}
	if err := c.Send(photo); err != nil {
		return err
	}
	// Читаем описание из файла
	description, err := readDescriptionFromFile(descriptionFilePath)
	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка при чтении описания: %v", err))
	}
	// Отправляем описание с клавиатурой
	return c.Send(description, keyboard)
}

// Функция для чтения описания из файла
func readDescriptionFromFile(filePath string) (string, error) {
	fullPath := filepath.Join("/home/pyssy/VSC/GO/bots/m_test_bot_2/pkg/txt", filePath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return "", fmt.Errorf("файл не найден: %s", fullPath)
	}

	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения файла: %v", err)
	}

	return string(data), nil
}

func messageHandler(c tele.Context) error {
	text := c.Text()
	fmt.Println("Получено сообщение:", text)

	// Карта для сопоставления текстовых команд с действиями
	actions := map[string]struct {
		message  string
		keyboard *tele.ReplyMarkup
		handler  func(c tele.Context) error
	}{
		// вываливающиеся менюшки
		"Продукты":         {"Меню продуктов", keyboard.CreateMeatProductMenu(), nil},
		"Меню":             {"Меню", keyboard.CreateMenu(), nil},
		"Мясо":             {"Мясо", keyboard.CreateMeatMenu(), nil},
		"Готовые продукты": {"Готовые продукты", keyboard.CreateReadyMealsMenu(), nil},
		"Акции": {"Акции", nil, func(c tele.Context) error {
			return handleCommand(c, "Акции", "говядина.jpg", "so.txt", nil)
		}},

		// главное меню
		"Заказ":    {"Здесь будут данные заказов", nil, nil},
		"Контакты": {"Контакты", nil, nil},
		"Доставка": {"Доставка", nil, nil},
		"Профиль":  {"Профиль", nil, handleInfoCommand},

		// меню сырого мяса
		"Говядина": {"Говядина", nil, func(c tele.Context) error {
			handleCommand(c, "Говядина", "говядина.jpg", "говядина.txt", keyboard.CreateKeyboard())
			return nil
		}},
		"Конина": {"Конина", nil, func(c tele.Context) error {
			return handleCommand(c, "Конина", "конина.jpg", "конина.txt", nil)
		}},
		"Курица": {"Курица", nil, func(c tele.Context) error {
			return handleCommand(c, "Курица", "курица.jpg", "курица.txt", nil)
		}},
		"Баранина": {"Баранина", nil, func(c tele.Context) error {
			return handleCommand(c, "Баранина", "баранина.jpg", "баранина.txt", nil)
		}},
		"Гусь": {"Гусь", nil, func(c tele.Context) error {
			return handleCommand(c, "Гусь", "гусь.jpg", "гусь.txt", nil)
		}},
		"Утка": {"Утка", nil, func(c tele.Context) error {
			return handleCommand(c, "Утка", "утка.jpg", "утка.txt", nil)
		}},

		// меню готовых продуктов
		"Казылык": {"Казылык", nil, nil},
		"Тур":     {"Тур", nil, nil},
	}

	if text == "Профиль" {
		c.Send("Ваши данные записаны!")
		return handleInfoCommand(c)
	}

	if action, exists := actions[text]; exists {
		if action.handler != nil {
			return action.handler(c)
		}
		if action.keyboard != nil {
			return c.Send(action.message, action.keyboard)
		}
		return c.Send(action.message)
	}

	return c.Send("Мы получили ваше сообщение и обрабатываем его!")
}

func RegisterCallback(bot *tele.Bot) {
	bot.Handle(&keyboard.Steak, func(c tele.Context) error {
		fmt.Println("Вы выбрали стейк!")
		return c.Send("Вы выбрали стейк!")
	})
	bot.Handle(&keyboard.Ribs, func(c tele.Context) error {
		fmt.Println("Вы выбрали ребра!")
		return c.Send("Вы выбрали ребра!")
	})
	bot.Handle(&keyboard.Mince, func(c tele.Context) error {
		fmt.Println("Вы выбрали фарш!")
		return c.Send("Вы выбрали фарш!")
	})
	bot.Handle(&keyboard.Ribeye, func(c tele.Context) error {
		fmt.Println("Вы выбрали рибай!")
		return c.Send("Вы выбрали рибай!")
	})
	bot.Handle(&keyboard.Liver, func(c tele.Context) error {
		fmt.Println("Вы выбрали печень!")
		return c.Send("Вы выбрали печень!")
	})
}
