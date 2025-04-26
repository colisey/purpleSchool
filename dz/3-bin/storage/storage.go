package storage

import (
	"encoding/json"
	"fmt"
	"main/bins"
	"main/files"
	"main/utils"
	"os"
)

// Однако, в пакете storage отсутствует работа с сущностями Bin и BinList,
// а реализованы только низкоуровневые функции записи и чтения файлов.
// Не реализованы функции сериализации/десериализации BinList,
// а также не обеспечена проверка расширения файла на json в storage.

// storage:
// - Сохранение bin в виде json в локальном файле
// - Чтение списка bin в виде json из локального файла

func WriteFile(binList *bins.BinList, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var content []byte
	isJSON := utils.IsJSONFile(name) // - Проверка что это json расширение файла err
	if isJSON {
		content, err = json.Marshal(binList)
		// return binList, nil
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}

func ReadFile(name string) (*bins.BinList, error) {
	data, err := files.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	isJSON := utils.IsJSONFile(name) // - Проверка что это json расширение файла err
	if isJSON {
		binList := bins.NewBinList(data)
		return binList, nil
	}
	return nil, nil
}
