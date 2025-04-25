package files

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func ReadFile(name string) ([]byte, error) { // - Чтение любого файла name
	isJSON := IsJSONFile(name) // - Проверка что это json расширение файла err
	if !isJSON {
		color.Red("Файл %s не является файлом *.json", name)
		return nil, errors.New("DONT_JSON")
	}

	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func IsJSONFile(filename string) bool {
	return strings.HasSuffix(filename, ".json")
}

func WriteFile(content []byte, name string) error {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Запись успешна")
	return nil
}
