package key

import (
	"bufio"
	"fmt"
	"os"
)

func ReadKey(filename string) (string, error) {
	// Открываем файл
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Создаём сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Читаем первую строку
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	// Проверяем на ошибки сканирования
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Если файл пустой
	return "", fmt.Errorf("file is empty")
}
