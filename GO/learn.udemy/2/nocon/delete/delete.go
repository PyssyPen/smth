package delete

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Delete() {
	a := ""
	fmt.Println("Hадо удалить логи? (y/n)")
	fmt.Fscanln(os.Stdin, &a)

	switch a {
	case "y", "Y", "Н", "н":
		fmt.Println("Удаляем логи...")
		// Укажите путь к папке, из которой нужно удалить файлы
		folderPath := "/home/pyssy/VSC/GO/learn.udemy/2/nocon/logs"

		// Используем filepath.WalkDir для обхода всех файлов в папке
		err := filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			// Проверяем, является ли файл .txt
			if !d.IsDir() && filepath.Ext(path) == ".txt" {
				// Удаляем файл
				err := os.Remove(path)
				if err != nil {
					return err
				}
				//println("Удален:", path)
			}
			return nil
		})

		if err != nil {
			println("Ошибка:", err)
		}
		println("Удалены")

	case "n", "N", "т", "Т":
		fmt.Println("Логи не удаляются.")
	default:
		fmt.Println("Некорректный ввод.")
	}

}
