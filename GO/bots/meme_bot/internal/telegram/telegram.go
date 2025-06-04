// package telegram

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"time"

// 	tele "gopkg.in/telebot.v4"
// )

// type memeData struct {
// 	PhotoPath  string
// 	TopText    string
// 	BottomText string
// }

// var storage = map[int64]memeData{}

// type ImageInterface interface {
// 	DrawText(inputFileName, topText, bottomText string) (string, error)
// }

// type Service struct {
// 	b              *tele.Bot
// 	imageInterface ImageInterface
// }

// func NewTelegramService(token string, imageInterface ImageInterface) *Service {
// 	pref := tele.Settings{
// 		Token:  token,
// 		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
// 	}

// 	b, err := tele.NewBot(pref)
// 	if err != nil {
// 		panic(err)
// 	}

// 	b.Handle(tele.OnPhoto, func(c tele.Context) error {
// 		photo := c.Message().Photo
// 		rc, err := b.File(&tele.File{FileID: photo.FileID})
// 		fmt.Println(rc) //
// 		if err != nil {
// 			return err
// 		}
// 		f, err := os.CreateTemp("", "telegramupload")
// 		fmt.Println(f) //
// 		if err != nil {
// 			return err
// 		}
// 		_, err = io.Copy(f, rc)
// 		if err != nil {
// 			return err
// 		}
// 		storage[c.Sender().ID] = memeData{
// 			PhotoPath: f.Name(),
// 		}
// 		fmt.Println(storage)
// 		//c.Send("Вот такая ширина ")
// 		//c.Send("Вот такая высота ")

// 		return c.Send("Теперь скинь верхний текст")
// 	})

// 	b.Handle(tele.OnText, func(c tele.Context) error {
// 		if memeData, ok := storage[c.Sender().ID]; ok {
// 			if memeData.TopText == "" {
// 				memeData.TopText = c.Message().Text
// 				storage[c.Sender().ID] = memeData
// 				return c.Send("Теперь скинь нижний текст")
// 			} else {
// 				result, err := imageInterface.DrawText(memeData.PhotoPath, memeData.TopText, c.Message().Text)
// 				if err != nil {
// 					return err
// 				}

// 				return c.Send(&tele.Photo{File: tele.FromDisk(result)})
// 			}
// 		}
// 		return nil
// 	})
// 	return &Service{b: b, imageInterface: imageInterface}
// }

// func (s *Service) Start() {
// 	s.b.Start()

// }

// func (s *Service) Stop() {
// 	s.b.Stop()
// }

package telegram

import (
	"io"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

type memeData struct {
	PhotoPath  string
	TopText    string
	BottomText string
}

var storage = map[int64]memeData{}

type ImageInterface interface {
	DrawText(inputFileName, topText, bottomText string) (string, error)
}

type Service struct {
	b              *tele.Bot
	imageInterface ImageInterface
}

func NewTelegramService(token string, imageInterface ImageInterface) *Service {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	b.Handle(tele.OnPhoto, func(c tele.Context) error {
		photo := c.Message().Photo
		rc, err := b.File(&tele.File{FileID: photo.FileID})
		if err != nil {
			return err
		}

		f, err := os.CreateTemp("/home/pyssy/VSC/GO/bots/meme_bot/jpg", "P_")
		if err != nil {
			return err
		}

		// Копируем содержимое файла в временный файл
		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}

		// Сохраняем путь к временному файлу в хранилище
		storage[c.Sender().ID] = memeData{
			PhotoPath: f.Name(),
		}

		return c.Send("Теперь скинь верхний текст")
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		if memeData, ok := storage[c.Sender().ID]; ok {
			if memeData.TopText == "" {
				memeData.TopText = c.Message().Text
				storage[c.Sender().ID] = memeData
				return c.Send("Теперь скинь нижний текст")
			} else {
				result, err := imageInterface.DrawText(memeData.PhotoPath, memeData.TopText, c.Message().Text)
				if err != nil {
					return err
				}

				return c.Send(&tele.Photo{File: tele.FromDisk(result)})
			}
		}
		return nil
	})
	return &Service{b: b, imageInterface: imageInterface}
}

func (s *Service) Start() {
	s.b.Start()
}

func (s *Service) Stop() {
	s.b.Stop()
}
