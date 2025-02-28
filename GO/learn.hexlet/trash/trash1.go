package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func IntToString(s int) string {
	return strconv.Itoa(s)
}

type Message struct {
	Sender string `json:"sender"` // ставим тег с описанием JSON поля
	Text   string `json:"text"`
}

// инициализация ошибки через конструктор стандартного пакета errors
var errEmptyMessage = errors.New("empty message")

// возвращаем ошибку в случае неожиданного поведения
func DecodeJSON(rawMsg string) (Message, error) {
	// если нам передали пустую строку, возвращаем ошибку об этом
	if len(rawMsg) == 0 {
		return Message{}, errEmptyMessage
	}

	msg := Message{}

	// декодируем строку в структуру
	err := json.Unmarshal([]byte(rawMsg), &msg)
	if err != nil {
		return Message{}, fmt.Errorf("unmarshal: %w", err)
	}

	return msg, nil
}
