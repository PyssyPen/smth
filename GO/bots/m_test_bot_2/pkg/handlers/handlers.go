package handlers

import (
	"fmt"
	"io/ioutil"
	"m_test_bot_2/pkg/cart"
	keyboard "m_test_bot_2/pkg/keyboards"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"

	tele "gopkg.in/telebot.v4"
)

var (
	waitingForKg1        = make(map[int64]bool)
	waitingForKg2        = make(map[int64]bool)
	ZeroString    string = ""
)

func logMessage(userID int64, text string) error {
	logDir := "/home/pyssy/VSC/GO/bots/m_test_bot_2/logs"
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

func handleInfoCommand(c tele.Context) error {
	user := c.Sender()
	if err := logMessage(user.ID, "Запрос профиля"); err != nil {
		fmt.Printf("Ошибка логирования: %v\n", err)
	}
	fmt.Printf("Запрос профиля UserID: %d\n", user.ID)
	return nil
}

func RegisterHandlers(bot *tele.Bot) {
	bot.Handle("/start", startHandler)
	bot.Handle("/cart", handleCartCommand)
	bot.Handle(tele.OnText, messageHandler)
}

func startHandler(c tele.Context) error {
	user := c.Sender()

	// Формируем строку с данными пользователя
	userInfo := fmt.Sprintf("User Info:\nUser ID: %d\nUsername: %s\nFirst Name: %s\nLast Name: %s\nLanguage Code: %s\nIs Bot: %v",
		user.ID, user.Username, user.FirstName, user.LastName, user.LanguageCode, user.IsBot)

	// Логируем данные пользователя
	if err := logMessage(user.ID, userInfo); err != nil {
		fmt.Printf("Ошибка логирования: %v\n", err)
	}

	if err := logMessage(user.ID, "/start"); err != nil {
		fmt.Printf("Ошибка логирования: %v\n", err)
	}
	fmt.Printf("Обработка /start для UserID: %d\n", user.ID)
	return c.Send(handleCommandStart(c, "start.txt", keyboard.CreateMenu()))
}

func handleCommandStart(c tele.Context, filePath string, keyboard *tele.ReplyMarkup) error {
	message, err := readDescriptionFromFile(filePath)
	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка при чтении файла: %v", err))
	}
	return c.Send(message, keyboard)
}

func handleCommand(c tele.Context, meatName, imagePath, descriptionFilePath string, keyboard *tele.ReplyMarkup) error {
	if err := logMessage(c.Sender().ID, fmt.Sprintf("Выбор: %s", meatName)); err != nil {
		fmt.Printf("Ошибка логирования: %v\n", err)
	}
	if err := c.Send(meatName); err != nil {
		return err
	}
	photo := &tele.Photo{File: tele.FromDisk(filepath.Join("/home/pyssy/VSC/GO/bots/m_test_bot_2/pkg/jpg", imagePath))}
	if err := c.Send(photo); err != nil {
		return err
	}
	description, err := readDescriptionFromFile(descriptionFilePath)
	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка при чтении описания: %v", err))
	}
	return c.Send(description, keyboard)
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
		fmt.Printf("Ошибка логирования: %v\n", err)
	}
	fmt.Printf("Получено сообщение: %s, UserID: %d\n", text, user.ID)

	if waitingForKg1[user.ID] {
		return KgHandler(ZeroString, c)
	}
	if waitingForKg2[user.ID] {
		return KgHandler(ZeroString, c)
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
		"Корзина":  {"Здесь будут данные о заказе, который формируется покупателем в данный момент", nil, nil},
		"Контакты": {"Здесь будут контакты продавца (и возможно курьера)", nil, nil},
		"Доставка": {"Здесь будет инфо либо о доставке уже заказанного набора продуктов, либо об доставке в общем", nil, nil},
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
		"Казылык": {"Казылык", nil, nil},
		"Тур":     {"Тур", nil, nil},
	}
	if text == "Мои данные" {
		c.Send("Ваши данные:")
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
	return c.Send("Мы получили ваше сообщение и обрабатываем его!", keyboard.CreateMenu())
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}

func extractNumber(input string) (float64, error) {
	var numberStr string
	numberStr = strings.ReplaceAll(numberStr, ",", ".")
	for _, char := range input {
		if unicode.IsDigit(char) || char == '.' {
			numberStr += string(char)
		} else if char == 'к' || char == 'К' || char == 'k' || char == 'K' {
			numberStr += string('.')
		}
	}
	if numberStr == "" {
		return 0, fmt.Errorf("no digits found in input")
	}
	return strconv.ParseFloat(numberStr, 64)
}

func KgHandler(meatType string, c tele.Context) error {
	if meatType == "" {
		return fmt.Errorf(meatType)
	}
	text := c.Text()
	userID := c.Sender().ID

	fmt.Printf("Обработка мяса: %s для UserID: %d\n", meatType, userID)

	if waitingForKg1[userID] {
		if kg, err := extractNumber(text); err == nil {
			fmt.Println(kg)
			if kg == 0 {
				delete(waitingForKg1, userID)
				return c.Send("Вы отменили выбор этого мяса.", keyboard.CreateMenu())
			} else if kg > 0 && kg < 15 {
				return c.Send("Мы доставляем от 15 кг!")
			}
			delete(waitingForKg1, userID)
			c.Send(fmt.Sprintf("Вы ввели %.1f килограмм. Правильно?", kg), keyboard.CreateYesNo())

			waitingForKg2[userID] = true
			strKg := strconv.FormatFloat(kg, 'f', -1, 64)
			logMessage(userID, strKg)

			// Добавление в корзину
			if err := cart.AddToCart(userID, meatType, kg); err != nil {
				return c.Send(fmt.Sprintf("Ошибка при добавлении в корзину: %v", err))
			}
			return nil
		} else {
			return c.Send("Пожалуйста, введите корректное число.")
		}
	}

	if waitingForKg2[userID] {
		affirmativesYes := []string{"Да", "да", "Yes", "yes", "+", "Правильно", "правильно"}
		denialNo := []string{"Нет", "нет", "No", "no", "-"}
		if contains(affirmativesYes, text) {
			delete(waitingForKg2, userID)
			return c.Send("Добавили в корзину!", keyboard.CreateMeatMenu())
		} else if contains(denialNo, text) {
			waitingForKg1[userID] = true
			return c.Send("Введите правильный вес:")
		}
	}

	if err := c.Send("Сколько килограмм мяса вам нужно? \n(если вы не хотели выбирать это мясо, то введите 0)"); err != nil {
		return err
	}
	waitingForKg1[userID] = true

	return nil
}

func RegisterCallback(bot *tele.Bot) {
	bot.Handle(&keyboard.Steak, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Вырезка"); err != nil {
			fmt.Printf("Ошибка логирования: %v\n", err)
		}
		fmt.Printf("Выбор вырезки UserID: %d\n", userID)
		ZeroString = "Вырезка"
		c.Send("Вы выбрали вырезку!")
		return KgHandler(ZeroString, c)
	})
	bot.Handle(&keyboard.Ribs, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Ребра"); err != nil {
			fmt.Printf("Ошибка логирования: %v\n", err)
		}
		fmt.Printf("Выбор ребер UserID: %d\n", userID)
		c.Send("Вы выбрали ребра!")
		return KgHandler("Ребра", c)
	})
	bot.Handle(&keyboard.Mince, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Фарш"); err != nil {
			fmt.Printf("Ошибка логирования: %v\n", err)
		}
		fmt.Printf("Выбор фарша UserID: %d\n", userID)
		c.Send("Вы выбрали фарш!")
		return KgHandler("Фарш", c)
	})
	bot.Handle(&keyboard.Ribeye, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Рибай"); err != nil {
			fmt.Printf("Ошибка логирования: %v\n", err)
		}
		fmt.Printf("Выбор рибай UserID: %d\n", userID)
		c.Send("Вы выбрали рибай!")
		return KgHandler("Рибай", c)
	})
	bot.Handle(&keyboard.Liver, func(c tele.Context) error {
		userID := c.Sender().ID
		if err := logMessage(userID, "Нажата кнопка: Печень"); err != nil {
			fmt.Printf("Ошибка логирования: %v\n", err)
		}
		fmt.Printf("Выбор печени UserID: %d\n", userID)
		c.Send("Вы выбрали печень!")
		return KgHandler("Печень", c)
	})
}

func handleCartCommand(c tele.Context) error {
	userID := c.Sender().ID
	cartItems, err := cart.GetCart(userID)
	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка при получении корзины: %v", err))
	}

	if len(cartItems) == 0 {
		return c.Send("Ваша корзина пуста.")
	}

	var message strings.Builder
	message.WriteString("Ваша корзина:\n")
	for product, quantity := range cartItems {
		message.WriteString(fmt.Sprintf("%s: %.1f кг\n", product, quantity))
	}
	return c.Send(message.String(), keyboard.CreateMenu())
}
