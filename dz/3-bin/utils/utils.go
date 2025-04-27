package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateProductID() string {
	currentTime := time.Now().UnixNano()
	randomNum := rand.Intn(1000) // случайное число
	return fmt.Sprintf("%d-%d", currentTime, randomNum)
}
func IsJSONFile(filename string) bool {
	return strings.HasSuffix(filename, ".json")
}

//	func ConvertToBytes(binList any, name string) ([]byte, error) {
//		var data []byte
//		var err error
//		isJSON := IsJSONFile(name) // - Проверка что это json расширение файла err
//		if isJSON {
//			data, err = json.Marshal(binList)
//			if err != nil {
//				return nil, err
//			}
//		} else {
//			return nil, errors.New("Неверный формат файла")
//		}
//		return data, nil
//	}
// func ConvertToData(data []byte, name string) (*bins.BinList, error) {
// 	var file bins.BinList
// 	var err error
// 	isJSON := IsJSONFile(name) // - Проверка что это json расширение файла err
// 	if isJSON {
// 		err = json.Unmarshal(data, &file)
// 		if err != nil {
// 			color.Red("Не удалось разобрать файл data.json")
// 			return nil, err
// 		}
// 		return &file, nil
// 	} else {
// 		return nil, errors.New("Неверный формат файла")
// 	}
// }
