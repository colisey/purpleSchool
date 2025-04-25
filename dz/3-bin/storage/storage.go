package storage

import (
	"fmt"
	"os"
)

// storage:
// - Сохранение bin в виде json в локальном файле
// - Чтение списка bin в виде json из локального файла

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
