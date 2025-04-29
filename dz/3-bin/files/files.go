package files

import (
	"encoding/json"
	"fmt"
	"main/bins"
	"main/utils"
	"os"
)

func ReadFile(name string) (*bins.BinList, error) { // - Чтение любого файла name
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var binList bins.BinList
	isJSON := utils.IsJSONFile(name) // - Проверка что это json расширение файла err
	if isJSON {
		err = json.Unmarshal(data, &binList)
		if err != nil {
			return nil, err
		}
	}

	return &binList, nil
}

func WriteFile(data *bins.BinList, name string) error {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	var content []byte
	isJSON := utils.IsJSONFile(name) // - Проверка что это json расширение файла err
	if isJSON {
		content, err = json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Запись успешна")
	return nil
}
