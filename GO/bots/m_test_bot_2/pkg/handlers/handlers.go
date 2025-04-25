package handlers

import (
	"fmt"
	"io/ioutil"
	"m_test_bot_2/pkg/cart"
	keyboard "m_test_bot_2/pkg/keyboards"
	"m_test_bot_2/pkg/profile"
	"os"
	"path/filepath"
	"time"

	tele "gopkg.in/telebot.v4"
)

var (
	waitingForKg1             = make(map[int64]bool)
	waitingForKg2             = make(map[int64]bool)
	waitingForUsername        = make(map[int64]bool)
	waitingForPhone           = make(map[int64]bool)
	waitingForAddress         = make(map[int64]bool)
	ZeroString         string = ""
	UserID             int64  = 0
)

const (
	logDir = "/home/pyssy/VSC/GO/bots/m_test_bot_2/logs"
)

func logMessage(userID int64, text string) error {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}
	logFile := filepath.Join(logDir, fmt.Sprintf("history_ID%d.txt", userID))
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s\n", timestamp, text)
	if _, err := file.WriteString(logEntry); err != nil {
		return err
	}
	return nil
}

func sendMessage(c tele.Context, text string, replyMarkup *tele.ReplyMarkup) error {
	return c.Send(text, replyMarkup)
}

func RegisterHandlers(bot *tele.Bot) {
	bot.Handle("/start", startHandler)
	bot.Handle(tele.OnText, messageHandler)
	RegisterCallback(bot)
}

func startHandler(c tele.Context) error {
	user := c.Sender()
	userInfo := fmt.Sprintf("User Info:\nUser ID: %d\nUsername: %s\nFirst Name: %s\nLast Name: %s\nLanguage Code: %s\nIs Bot: %v",
		user.ID, user.Username, user.FirstName, user.LastName, user.LanguageCode, user.IsBot)
	UserID = user.ID

	if err := logMessage(user.ID, userInfo); err != nil {
		return fmt.Errorf("ошибка логирования: %v", err)
	}

	if err := logMessage(user.ID, "/start"); err != nil {
		return fmt.Errorf("ошибка логирования: %v", err)
	}

	return c.Send(handleCommandStart(c, "start.txt", keyboard.CreateMenu()))
}

func handleCommandStart(c tele.Context, filePath string, keyboard *tele.ReplyMarkup) error {
	message, err := readDescriptionFromFile(filePath)
	if err != nil {
		return sendMessage(c, fmt.Sprintf("Ошибка при чтении файла: %v", err), nil)
	}
	return sendMessage(c, message, keyboard)
}

func handleCommand(c tele.Context, meatName, imagePath, descriptionFilePath string, keyboard *tele.ReplyMarkup) error {
	if err := logMessage(c.Sender().ID, fmt.Sprintf("Выбор: %s", meatName)); err != nil {
		return fmt.Errorf("ошибка логирования: %v", err)
	}
	if err := sendMessage(c, meatName, nil); err != nil {
		return err
	}
	photo := &tele.Photo{File: tele.FromDisk(filepath.Join("/home/pyssy/VSC/GO/bots/m_test_bot_2/pkg/jpg", imagePath))}
	if err := c.Send(photo); err != nil {
		return err
	}
	description, err := readDescriptionFromFile(descriptionFilePath)
	if err != nil {
		return sendMessage(c, fmt.Sprintf("Ошибка при чтении описания: %v", err), nil)
	}
	return sendMessage(c, description, keyboard)
}

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
	user := c.Sender()
	if err := logMessage(user.ID, text); err != nil {
		return fmt.Errorf("ошибка логирования: %v", err)
	}

	if waitingForKg1[user.ID] {
		return KgHandler(ZeroString, c)
	}
	if waitingForKg2[user.ID] {
		return KgHandler(ZeroString, c)
	}
	if waitingForUsername[user.ID] {
		return updateProfileHandler(c, "username", text)
	}
	if waitingForPhone[user.ID] {
		return updateProfileHandler(c, "phone_number", text)
	}
	if waitingForAddress[user.ID] {
		return updateProfileHandler(c, "delivery_address", text)
	}

	actions := map[string]struct {
		message  string
		keyboard *tele.ReplyMarkup
		handler  func(c tele.Context) error
	}{
		"Продукты":         {"Меню продуктов", keyboard.CreateMeatProductMenu(), nil},
		"Меню":             {"Меню", keyboard.CreateMenu(), nil},
		"Мясо":             {"Мясо", keyboard.CreateMeatMenu(), nil},
		"Готовые продукты": {"Готовые продукты", keyboard.CreateReadyMealsMenu(), nil},
		"Акции": {"Акции", nil, func(c tele.Context) error {
			return handleCommand(c, "Акции", "говядина.jpg", "so.txt", nil)
		}},
		"Контакты": {"По вопросам бота и заказа @penovich4 \nПо вопросам доставки @guru2341d", nil, nil},

		"Корзина":     {"Корзина", nil, handleCartCommand},
		"Мои заказы":  {"Заказ", nil, handleOrderCommand},
		"Мой профиль": {"Мой профиль", nil, handleProfileCommand},

		"Говядина": {"Говядина", nil, func(c tele.Context) error {
			handleCommand(c, "Говядина", "говядина.jpg", "говядина.txt", keyboard.CreateKeyboard())
			return nil
		}},
		"Конина": {"Конина", nil, func(c tele.Context) error {
			return handleCommand(c, "Конина", "конина.jpg", "конина.txt", keyboard.CreateHorseMeat())
		}},
		"Курица": {"Курица", nil, func(c tele.Context) error {
			return handleCommand(c, "Курица", "курица.jpg", "курица.txt", keyboard.CreateChicken())
		}},
		"Баранина": {"Баранина", nil, func(c tele.Context) error {
			return handleCommand(c, "Баранина", "баранина.jpg", "баранина.txt", keyboard.CreateLamb())
		}},
		"Гусь": {"Гусь", nil, func(c tele.Context) error {
			return handleCommand(c, "Гусь", "гусь.jpg", "гусь.txt", keyboard.CreateGoose())
		}},
		"Утка": {"Утка", nil, func(c tele.Context) error {
			return handleCommand(c, "Утка", "утка.jpg", "утка.txt", keyboard.CreateDuck())
		}},
		"Казылык": {"Казылык", nil, func(c tele.Context) error {
			return handleCommand(c, "Казылык", "казылык.jpg", "казылык.txt", keyboard.CreateKazylyk())
		}},
		"Тутырма": {"Тутырма", nil, func(c tele.Context) error {
			return handleCommand(c, "Тутырма", "тутырма.jpg", "тутырма.txt", keyboard.CreateTutyrma())
		}},
		"Изменить имя пользователя": {"Введите новое имя пользователя:", nil, func(c tele.Context) error {
			waitingForUsername[c.Sender().ID] = true
			return nil
		}},
		"Изменить номер телефона": {"Введите новый номер телефона:", nil, func(c tele.Context) error {
			waitingForPhone[c.Sender().ID] = true
			return nil
		}},
		"Изменить адрес доставки": {"Введите новый адрес доставки:", nil, func(c tele.Context) error {
			waitingForAddress[c.Sender().ID] = true
			return nil
		}},
	}

	if text == "Заказать" {
		userID := c.Sender().ID
		profileData, err := profile.GetProfile(userID)
		if err != nil {
			return sendMessage(c, fmt.Sprintf("Ошибка при получении профиля: %v", err), nil)
		}

		if profileData["username"] == "" || profileData["phone_number"] == "" || profileData["delivery_address"] == "" {
			return sendMessage(c, "Пожалуйста, заполните ваш профиль перед оформлением заказа. Нажмите 'Мой профиль' для заполнения данных.", nil)
		}

		if err := cart.CreateOrder(userID); err != nil {
			return sendMessage(c, fmt.Sprintf("Ошибка при создании заказа: %v", err), nil)
		}
		return sendMessage(c, "Ваш заказ успешно оформлен!", keyboard.CreateMenu())
	}

	if text == "Очистить" {
		userID := c.Sender().ID
		if err := cart.ClearCart(userID); err != nil {
			return sendMessage(c, fmt.Sprintf("Ошибка при очистке корзины: %v", err), nil)
		}
		return sendMessage(c, "Теперь корзина пуста!", keyboard.CreateMenu())
	}

	if action, exists := actions[text]; exists {
		if action.handler != nil {
			return action.handler(c)
		}
		if action.keyboard != nil {
			return sendMessage(c, action.message, action.keyboard)
		}
		return sendMessage(c, action.message, nil)
	}

	return sendMessage(c, "Мы получили ваше сообщение и обрабатываем его!", keyboard.CreateMenu())
}
