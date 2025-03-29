package keyboard

import (
	tele "gopkg.in/telebot.v4"
)

// кнопки в сообщении
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
			{
				{Text: "Меню"},
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
			{
				{Text: "Меню"},
			},
		},
	}
}

func CreateMeatMenu() *tele.ReplyMarkup {
	turkey := tele.InlineButton{
		Unique: "meat5",
		Text:   "Индейка",
	}
	chicken := tele.InlineButton{
		Unique: "meat6",
		Text:   "Курица",
	}
	beef := tele.InlineButton{
		Unique: "meat7",
		Text:   "Говядина",
	}
	contact := tele.InlineButton{
		Unique: "contact",
		Text:   "Контакты",
	}
	location := tele.InlineButton{
		Unique: "location",
		Text:   "Местоположение",
	}

	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{turkey, chicken, beef},
			{contact, location},
		},
	}
}

// func CreateMeatMenu() *tele.ReplyMarkup {
// 	return &tele.ReplyMarkup{
// 		ResizeKeyboard: true,
// 		ReplyKeyboard: [][]tele.ReplyButton{
// 			{
// 				{Text: "Говядина"},
// 				{Text: "Баранина"},
// 			},
// 			{
// 				{Text: "Конина"},
// 				{Text: "Конина"},
// 			},
// 		},
// 	}

// 	// beef := tele.InlineButton{
// 	// 	Unique: "meat1",
// 	// 	Text:   "Говядина",
// 	// }

// 	// // Добавьте остальные кнопки аналогично
// 	// pork := tele.InlineButton{
// 	// 	Unique: "meat2",
// 	// 	Text:   "Свинина",
// 	// }

// 	// chicken := tele.InlineButton{
// 	// 	Unique: "meat3",
// 	// 	Text:   "Курица",
// 	// }

// 	// lamb := tele.InlineButton{
// 	// 	Unique: "meat4",
// 	// 	Text:   "Баранина",
// 	// }

// 	// turkey := tele.InlineButton{
// 	// 	Unique: "meat5",
// 	// 	Text:   "Индейка",
// 	// }

// 	// duck := tele.InlineButton{
// 	// 	Unique: "meat6",
// 	// 	Text:   "Утка",
// 	// }

// 	// // Создайте и верните клавиатуру с этими кнопками
// 	// return &tele.ReplyMarkup{
// 	// 	InlineKeyboard: [][]tele.InlineButton{
// 	// 		{pork, beef},
// 	// 		{chicken, lamb},
// 	// 		{turkey, duck},
// 	// 	},
// 	// }
// }

func CreateReadyMealsMenu() *tele.ReplyMarkup {
	dumplings := tele.InlineButton{
		Unique: "dumplings",
		Text:   "Пельмени",
	}

	cutlets := tele.InlineButton{
		Unique: "cutlets",
		Text:   "Котлеты",
	}

	shashlik := tele.InlineButton{
		Unique: "shashlik",
		Text:   "Шашлык",
	}

	// Создайте и верните клавиатуру с этими кнопками
	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{dumplings, cutlets},
			{shashlik},
		},
	}
}
