package keyboard

import (
	tele "gopkg.in/telebot.v4"
)

// кнопки внизу
func CreateMenu() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{
				{Text: "Продукты"},
				{Text: "Акции"},
			},
			{
				{Text: "Заказ"},
				{Text: "Контакты"},
			},
			{
				{Text: "Доставка"},
				{Text: "Профиль"},
			},
		},
	}
}

func CreateMeatProductMenu() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{
				{Text: "Мясо"},
				{Text: "Готовые продукты"},
			},
			{{Text: "Меню"}},
		},
	}
}

func CreateReadyMealsMenu() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{
				{Text: "Казылык"},
				{Text: "Тур"},
			},
			{{Text: "Меню"}},
		},
	}
}

func CreateMeatMenu() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{
				{Text: "Говядина"},
				{Text: "Баранина"},
			},
			{
				{Text: "Конина"},
				{Text: "Гусь"},
			},
			{
				{Text: "Курица"},
				{Text: "Утка"},
			},
			{{Text: "Меню"}},
		},
	}
}

// Создаем кнопки в сообщении
var (
	Steak = tele.InlineButton{
		Unique: "Steak",
		Text:   "🥩 Вырезка",
	}
	Ribs = tele.InlineButton{
		Unique: "Ribs",
		Text:   "🍖 Ребра",
	}
	Mince = tele.InlineButton{
		Unique: "Mince",
		Text:   "🧆 Фарш",
	}
	Ribeye = tele.InlineButton{
		Unique: "Ribeye",
		Text:   "🌟 Рибай",
	}
	Liver = tele.InlineButton{
		Unique: "Liver",
		Text:   "💊 Печень",
	}

	Plus = tele.InlineButton{
		Unique: "plus",
		Text:   "+",
	}
	Minus = tele.InlineButton{
		Unique: "minus",
		Text:   "-",
	}
	Count = tele.InlineButton{
		Unique: "count",
		Text:   "1", // Начальное значение
	}
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
