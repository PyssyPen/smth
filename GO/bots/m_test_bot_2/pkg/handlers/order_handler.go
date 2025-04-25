package handlers

import (
	"fmt"
	keyboard "m_test_bot_2/pkg/keyboards"
	"m_test_bot_2/pkg/order"
	"strings"

	tele "gopkg.in/telebot.v4"
)

func handleOrderCommand(c tele.Context) error {
	userID := c.Sender().ID
	orderItems, err := order.GetOrders(userID)
	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка при получении заказов: %v", err))
	}

	if len(orderItems) == 0 {
		return c.Send("У вас не было заказов.")
	}

	var message strings.Builder
	message.WriteString("Ваши заказанные продукты:\n")
	for product, quantity := range orderItems {
		message.WriteString(fmt.Sprintf("%s: %.1f кг\n", product, quantity))
	}
	return c.Send(message.String(), keyboard.CreateMenu())
}
