package main

import (
	"fmt"
	"log"
	"m_test_bot/bot/bot"
	"m_test_bot/key"
)

func main() {
	// Получаем токен бота из переменной окружения
	token, err := key.ReadKey("/home/pyssy/VSC/GO/bots/m_test_bot/key/key.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = bot.StartBot(token)
	if err != nil {
		log.Fatal("Ошибка запуска бота:", err)
	}
}

/*
КНОПКИ СНИЗУ

сосиски убрать
свинину убрать
конину поставить
гуся добавить
индейку убрать
пельмени котлеты шашлык убрать

в готовые продукты казылык
казый поставить отдельно

подждержку можно поставит
*/

// import (
// 	"fmt"
// 	"log"
// 	"m_test_bot/key"
// 	"time"

// 	"gopkg.in/tucnak/telebot.v2"
// )

// func main() {
// 	// Получаем токен бота из переменной окружения
// 	token, err := key.ReadKey("/home/pyssy/VSC/GO/m_test_bot/key/key.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	bot, err := telebot.NewBot(telebot.Settings{
// 		Token:  token,
// 		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	bot.Handle(telebot.OnText, func(m *telebot.Message) {
// 		switch m.Text {
// 		case "info":
// 			user := m.Sender
// 			fmt.Printf("User Info:\n")
// 			fmt.Printf("User ID: %d\n", user.ID)
// 			fmt.Printf("Username: %s\n", user.Username)
// 			fmt.Printf("First Name: %s\n", user.FirstName)
// 			fmt.Printf("Last Name: %s\n", user.LastName)
// 			fmt.Printf("Language Code: %s\n", user.LanguageCode)
// 			fmt.Printf("Is Bot: %v\n", user.IsBot)

// 			bot.Send(m.Sender, "Ваша информация была получена!")
// 		}
// 	})

// 	bot.Start()
// }
