package handlers

import (
	"fmt"
	"strings"

	//"log"

	//"m_test_bot/bot/file"
	"m_test_bot/bot/keyboard"

	tele "gopkg.in/telebot.v4"
)

func RegisterHandlers(bot *tele.Bot) {
	bot.Handle("/start", startHandler)
	//bot.Handle("/help", helpHandler)
	bot.Handle(tele.OnText, messageHandler)
	bot.Handle(tele.OnCallback, callbackHandler)
	//bot.Handle("/buttons", buttonsHandler)
}

func startHandler(c tele.Context) error {
	fmt.Println("Получена команда /start")
	return c.Send("Привет! \nЯ бот магазина K’RASNAYA NEFT\nВнизу появилось меню с командами", keyboard.CreateMenu())
}

// тут выводим нужное меню из нижних кнопок
// Получено сообщение: Привет
func messageHandler(c tele.Context) error {
	text := c.Text()
	fmt.Println("Получено сообщение:", text)

	//return c.Send("КНОПКА СНИЗУ: "+text, keyboard.CreateMainMenu())
	//return c.Send("Меню:", keyboard.CreateMenu())
	switch text {
	case "Продукты":
		fmt.Println("отправляем клавиатуру Продукты ")
		c.Send("cghjklgckhgfc", keyboard.CreateMeatProductMenu())
		return nil

	case "Меню":
		return c.Send("орпм", keyboard.CreateMenu())

	// main menu
	case "products_text":
		return c.Send("Меню продуктов", keyboard.CreateMeatProductMenu())
	case "so_text":
		return c.Send("Меню акций") // keyboard.SO()
	case "orders_text":
		return c.Send("Здесь будут данные заказов") // keyboard.Orders()
	case "contacts_text":
		return c.Send("Контакты")
	case "delivery_text":
		return c.Send("Доставка")

	// meat product categories
	case "meat":
		return c.Send("Мясо", keyboard.CreateMeatMenu())
	// case "sausages":
	// 	return c.Send("Сосиски", keyboard.CreateSausageMenu())
	case "ready":
		return c.Send("Готовые продукты", keyboard.CreateReadyMealsMenu())

	// meat versions
	case "meat1":
		return c.Send("Говядина")
	case "meat2":
		return c.Send("Свинина")
	case "meat3":
		return c.Send("Курица")
	case "meat4":
		return c.Send("Баранина")
	case "meat5":
		return c.Send("Индейка")
	case "meat6":
		return c.Send("Утка")

	// sausage versions
	case "boiled_sausage":
		return c.Send("Вареные колбасы")
	case "semi_smoked_sausage":
		return c.Send("Полукопчёные колбасы")
	case "smoked_sausage":
		return c.Send("Сырокопчёные колбасы")

	// ready meals versions
	case "dumplings":
		return c.Send("Пельмени")
	case "cutlets":
		return c.Send("Котлеты")
	case "shashlik":
		return c.Send("Шашлык")

	default:
		return c.Send("Неизвестная кнопка!")
	}
}

// // тут выводим то, что было нажато в сообщении в виде кнопки
// // Получен callback: btn_1
func callbackHandler(c tele.Context) error {
	callback := c.Callback()
	if callback == nil {
		fmt.Println("Callback is nil")
		return c.Send("Ошибка обработки callback-запроса.")
	}

	callbackData := strings.TrimSpace(callback.Data)
	fmt.Println("Получен callback:", callbackData)

	switch callbackData {

	// case "Start", "start", "привет", "Привет":
	// 	return c.Send("Привет! \nЯ бот магазина K’RASNAYA NEFT\nВнизу появилось меню с командами", keyboard.CreateMainMenu())

	// case "info":
	// 	user := m.Sender
	// 	fmt.Printf("User Info:\n")
	// 	fmt.Printf("User ID: %d\n", user.ID)
	// 	fmt.Printf("Username: %s\n", user.Username)
	// 	fmt.Printf("First Name: %s\n", user.FirstName)
	// 	fmt.Printf("Last Name: %s\n", user.LastName)
	// 	fmt.Printf("Language Code: %s\n", user.LanguageCode)
	// 	fmt.Printf("Is Bot: %v\n", user.IsBot)
	// 	c.Send("Ваша информация была получена!")

	case "Продукты":
		return c.Send(keyboard.CreateMeatProductMenu())
	// main menu
	case "products_text":
		return c.Send("Меню продуктов", keyboard.CreateMeatProductMenu())
	case "so_text":
		return c.Send("Меню акций") // keyboard.SO()
	case "orders_text":
		return c.Send("Здесь будут данные заказов") // keyboard.Orders()
	case "contacts_text":
		return c.Send("Контакты")
	case "delivery_text":
		return c.Send("Доставка")

	// meat product categories
	case "meat":
		return c.Send("Мясо", keyboard.CreateMeatMenu())
	// case "sausages":
	// 	return c.Send("Сосиски", keyboard.CreateSausageMenu())
	case "ready":
		return c.Send("Готовые продукты", keyboard.CreateReadyMealsMenu())

	// meat versions
	case "meat1":
		return c.Send("Говядина")
	case "meat2":
		return c.Send("Свинина") // delete
	case "meat3":
		return c.Send("Курица")
	case "meat4":
		return c.Send("Баранина")
	case "meat5":
		return c.Send("Индейка") // replace goose
	case "meat6":
		return c.Send("Утка")

	// ready meals versions
	case "dumplings":
		return c.Send("Пельмени") // delete
	case "cutlets":
		return c.Send("Котлеты") // delete
	case "shashlik":
		return c.Send("Шашлык") // delete
	//казылык

	default:
		return c.Send("Неизвестная кнопка!")
	}
}
