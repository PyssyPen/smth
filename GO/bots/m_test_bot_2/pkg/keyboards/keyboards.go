package keyboard

import tele "gopkg.in/telebot.v4"

func CreateMenu() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{{Text: "Продукты"}, {Text: "Акции"}},
			{{Text: "Корзина"}, {Text: "Контакты"}},
			{{Text: "Доставка"}, {Text: "Мои данные"}},
		},
	}
}
func CreateMeatProductMenu() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{{Text: "Мясо"}, {Text: "Готовые продукты"}},
			{{Text: "Меню"}},
		},
	}
}
func CreateReadyMealsMenu() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{{Text: "Казылык"}, {Text: "Тур"}},
			{{Text: "Меню"}},
		},
	}
}
func CreateMeatMenu() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{{Text: "Говядина"}, {Text: "Баранина"}},
			{{Text: "Конина"}, {Text: "Гусь"}},
			{{Text: "Курица"}, {Text: "Утка"}},
			{{Text: "Меню"}},
		},
	}
}

var (
	Steak  = tele.InlineButton{Unique: "Steak", Text: "🥩 Вырезка"}
	Ribs   = tele.InlineButton{Unique: "Ribs", Text: "🍖 Ребра"}
	Mince  = tele.InlineButton{Unique: "Mince", Text: "🧆 Фарш"}
	Ribeye = tele.InlineButton{Unique: "Ribeye", Text: "🌟 Рибай"}
	Liver  = tele.InlineButton{Unique: "Liver", Text: "💊 Печень"}
	Plus   = tele.InlineButton{Unique: "plus", Text: "+"}
	Minus  = tele.InlineButton{Unique: "minus", Text: "-"}
	Count  = tele.InlineButton{Unique: "count", Text: "1"}
)

func CreateKeyboard() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{Steak, Ribs},
			{Mince, Ribeye},
			{Liver},
		},
	}
}

func CreateYesNo() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{{Text: "Да"}, {Text: "Нет"}},
		},
	}
}
