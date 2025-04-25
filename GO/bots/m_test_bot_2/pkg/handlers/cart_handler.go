package handlers

import (
	"fmt"
	"m_test_bot_2/pkg/cart"
	keyboard "m_test_bot_2/pkg/keyboards"
	"strings"

	tele "gopkg.in/telebot.v4"
)

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
	return c.Send(message.String(), keyboard.CreateCartMenu())
}
