package file

// сюда надо вставить логирование из какого-то урочка юдеми
// пока файл создается, но ничего не пишется

import (
	"fmt"
	"os"
)

func CreateHistoryFile(fileIndex int) (*os.File, int, error) {
	fileIndex++
	fileName := fmt.Sprintf("%d.txt", fileIndex)
	file, err := os.Create(fileName)
	if err != nil {
		return nil, fileIndex, err
	}
	return file, fileIndex, nil
}

func DeleteHistoryFile(fileIndex int) error {
	fileName := fmt.Sprintf("%d.txt", fileIndex)
	return os.Remove(fileName)
}
