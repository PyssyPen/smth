package handlers

import (
	"fmt"
	keyboard "m_test_bot_2/pkg/keyboards"
	"m_test_bot_2/pkg/profile"
	"strings"

	tele "gopkg.in/telebot.v4"
)

func handleProfileCommand(c tele.Context) error {
	userID := c.Sender().ID
	profileData, err := profile.GetProfile(userID)
	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка при получении профиля: %v", err))
	}

	var message strings.Builder
	message.WriteString("Ваш профиль:\n")
	message.WriteString(fmt.Sprintf("Имя: %s\n", profileData["username"]))
	message.WriteString(fmt.Sprintf("Номер телефона: %s\n", profileData["phone_number"]))
	message.WriteString(fmt.Sprintf("Адрес доставки: %s\n", profileData["delivery_address"]))

	return c.Send(message.String(), keyboard.CreateProfileMenu())
}

func updateProfileHandler(c tele.Context, field, value string) error {
	userID := c.Sender().ID

	switch field {
	case "username":
		c.Send("Введите новое имя:")
		delete(waitingForUsername, userID)
	case "phone_number":
		delete(waitingForPhone, userID)
	case "delivery_address":
		delete(waitingForAddress, userID)
	}

	if err := profile.UpdateProfile(userID, value, "", ""); err != nil {
		return c.Send(fmt.Sprintf("Ошибка при обновлении профиля: %v", err))
	}

	return c.Send("Ваш профиль успешно обновлен!", keyboard.CreateProfileMenu())
}
