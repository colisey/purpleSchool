package storage

import (
	"fmt"
	"main/bins"
	"main/files"
)

// Однако, в пакете storage отсутствует работа с сущностями Bin и BinList,
// а реализованы только низкоуровневые функции записи и чтения файлов.
// Не реализованы функции сериализации/десериализации BinList,
// а также не обеспечена проверка расширения файла на json в storage.

// storage:
// - Сохранение bin в виде json в локальном файле
// - Чтение списка bin в виде json из локального файла

func Write(binList *bins.BinList, name string) {
	// var content []byte
	// var err error
	// isJSON := utils.IsJSONFile(name) // - Проверка что это json расширение файла err
	// if isJSON {
	// 	content, err = json.Marshal(binList)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }
	err := files.Write(binList, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}

func Read(name string) (*bins.BinList, error) {
	data, err := files.Read(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
