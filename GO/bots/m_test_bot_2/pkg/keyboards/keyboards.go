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
			{{Text: "Тутырма"}},
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
	Steak      = tele.InlineButton{Unique: "Steak", Text: "Вырезка"}
	Ribs       = tele.InlineButton{Unique: "Ribs", Text: "Ребра"}
	Mince      = tele.InlineButton{Unique: "Mince", Text: "Фарш"}
	Ribeye     = tele.InlineButton{Unique: "Ribeye", Text: "Рибай"}
	Liver      = tele.InlineButton{Unique: "Liver", Text: "Печень"}
	FrontThigh = tele.InlineButton{Unique: "FrontThigh", Text: "Ляжка передняя"}
	RearThigh  = tele.InlineButton{Unique: "RearThigh", Text: "Ляжка задняя"}

	Lamb      = tele.InlineButton{Unique: "Lamb", Text: "Баранина"}
	Goose     = tele.InlineButton{Unique: "Goose", Text: "Гусь"}
	Duck      = tele.InlineButton{Unique: "Duck", Text: "Утка"}
	HorseMeat = tele.InlineButton{Unique: "HorseMeat", Text: "Конина"}
	Chicken   = tele.InlineButton{Unique: "Chicken", Text: "Курица"}
	Kazylyk   = tele.InlineButton{Unique: "Kazylyk", Text: "Казылык"}
	Turkey    = tele.InlineButton{Unique: "Turkey", Text: "Тур"}
	Tutyrma   = tele.InlineButton{Unique: "Tutyrma", Text: "Тутырма"}
)

func CreateKeyboard() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{FrontThigh, RearThigh},
			{Steak, Ribs},
			{Mince, Ribeye},
			{Liver},
		},
	}
}

func CreateLamb() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{Lamb},
		},
	}
}

func CreateGoose() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{Goose},
		},
	}
}

func CreateDuck() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{Duck},
		},
	}
}

func CreateHorseMeat() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{HorseMeat},
		},
	}
}

func CreateChicken() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{Chicken},
		},
	}
}

func CreateKazylyk() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{Kazylyk},
		},
	}
}

func CreateTurkey() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{Turkey},
		},
	}
}

func CreateTutyrma() *tele.ReplyMarkup {
	return &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{Tutyrma},
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
