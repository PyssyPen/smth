package handlers

import (
	"fmt"
	"m_test_bot_2/pkg/cart"
	keyboard "m_test_bot_2/pkg/keyboards"
	"strconv"
	"strings"
	"unicode"

	tele "gopkg.in/telebot.v4"
)

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}

func extractNumber(input string) (float64, error) {
	var numberStr string
	numberStr = strings.ReplaceAll(numberStr, ",", ".")
	for _, char := range input {
		if unicode.IsDigit(char) || char == '.' {
			numberStr += string(char)
		} else if char == 'к' || char == 'К' || char == 'k' || char == 'K' {
			numberStr += string('.')
		}
	}
	if numberStr == "" {
		return 0, fmt.Errorf("no digits found in input")
	}
	return strconv.ParseFloat(numberStr, 64)
}

func KgHandler(meatType string, c tele.Context) error {
	text := c.Text()
	userID := c.Sender().ID

	if waitingForKg1[userID] {
		if kg, err := extractNumber(text); err == nil {
			if kg == 0 {
				delete(waitingForKg1, userID)
				return sendMessage(c, "Вы отменили выбор этого мяса.", keyboard.CreateMenu())
			} else if kg > 0 && kg < 15 {
				return sendMessage(c, "Мы доставляем от 15 кг!", nil)
			}
			delete(waitingForKg1, userID)
			sendMessage(c, fmt.Sprintf("Вы ввели %.1f килограмм. Правильно?", kg), keyboard.CreateYesNo())

			waitingForKg2[userID] = true
			strKg := strconv.FormatFloat(kg, 'f', -1, 64)
			if err := logMessage(userID, strKg); err != nil {
				return fmt.Errorf("ошибка логирования: %v", err)
			}

			if err := cart.AddToCart(userID, meatType, kg); err != nil {
				return sendMessage(c, fmt.Sprintf("Ошибка при добавлении в корзину: %v", err), nil)
			}
			return nil
		} else {
			return sendMessage(c, "Пожалуйста, введите корректное число.", nil)
		}
	}

	if waitingForKg2[userID] {
		affirmativesYes := []string{"Да", "да", "Yes", "yes", "+", "Правильно", "правильно"}
		denialNo := []string{"Нет", "нет", "No", "no", "-"}
		if contains(affirmativesYes, text) {
			delete(waitingForKg2, userID)
			return sendMessage(c, "Добавили в корзину!", keyboard.CreateMeatMenu())
		} else if contains(denialNo, text) {
			waitingForKg1[userID] = true
			return sendMessage(c, "Введите правильный вес:", nil)
		}
	}

	if err := sendMessage(c, "Сколько килограмм мяса вам нужно? \n(если вы не хотели выбирать это мясо, то введите 0)", nil); err != nil {
		return err
	}
	waitingForKg1[userID] = true

	return nil
}
