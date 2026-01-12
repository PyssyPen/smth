package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден")
	}
}

func main() {
	token := os.Getenv("TOKEN")

	go func() {
		if err := StartBot(token); err != nil {
			log.Fatalf("Ошибка запуска бота: %v", err)
		}
	}()

	select {}
}

func StartBot(token string) error {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return err
	}

	btnSong := telebot.InlineButton{Unique: "song", Text: "Какая песня"}
	btnBack := telebot.InlineButton{Unique: "back", Text: "Назад"}
	btnPlayPause := telebot.InlineButton{Unique: "playpause", Text: "Старт / Пауза"}
	btnNext := telebot.InlineButton{Unique: "next", Text: "Вперед"}
	btnPlus := telebot.InlineButton{Unique: "plus", Text: "+"}
	btnMinus := telebot.InlineButton{Unique: "minus", Text: "-"}

	keyboard := [][]telebot.InlineButton{
		{btnSong},
		{btnBack, btnPlayPause, btnNext},
		{btnMinus, btnPlus},
	}

	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Меню управления музыкой:", &telebot.ReplyMarkup{
			InlineKeyboard: keyboard,
		})
	})

	bot.Handle(&btnSong, func(c telebot.Context) error {
		song := playerctlSong()
		c.Edit("Сейчас играет: " + song)
		return c.Send("Меню управления музыкой:", &telebot.ReplyMarkup{
			InlineKeyboard: keyboard,
		})
	})

	bot.Handle(&btnBack, func(c telebot.Context) error {
		playerctlPrev()
		return c.Respond()
	})

	bot.Handle(&btnPlayPause, func(c telebot.Context) error {
		playerctlPlayPause()
		return c.Respond()
	})

	bot.Handle(&btnNext, func(c telebot.Context) error {
		playerctlNext()
		return c.Respond()
	})

	bot.Handle(&btnPlus, func(c telebot.Context) error {
		pactlVolumeUp()
		return c.Respond()
	})

	bot.Handle(&btnMinus, func(c telebot.Context) error {
		pactlVolumeDown()
		return c.Respond()
	})

	fmt.Println("Бот стартовал!")
	bot.Start()
	return nil
}

// pactl get-sink-volume @DEFAULT_SINK@ | grep -oP '\d+%' | head -1
func pactlVolumeUp() {
	cmd := exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", "+5%")
	err := cmd.Run()
	if err != nil {
		log.Println("Ошибка pactlVolumeUp:", err)
	}
}

func pactlVolumeDown() {
	cmd := exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", "-5%")
	err := cmd.Run()
	if err != nil {
		log.Println("Ошибка pactlVolumeDown:", err)
	}
}

func playerctlSong() string {
	cmd := exec.Command("playerctl", "metadata", "--format", "{{artist}} - {{title}}")
	output, err := cmd.Output()
	if err != nil {
		log.Println("Ошибка playerctlSong:", err)
		return "Информация недоступна"
	}
	return string(output)
}

func playerctlNext() {
	cmd := exec.Command("playerctl", "next")
	err := cmd.Run()
	if err != nil {
		log.Println("Ошибка playerctlNext:", err)
	}
}

func playerctlPrev() {
	cmd := exec.Command("playerctl", "previous")
	err := cmd.Run()
	if err != nil {
		log.Println("Ошибка playerctlPrev:", err)
	}
}

func playerctlPlayPause() {
	err := exec.Command("playerctl", "play-pause").Run()
	if err != nil {
		log.Println("Ошибка playerctlPlayPause:", err)
	}
}
