package utils

import (
	"encoding/json"
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
func ConvertToBytes(binList any, name string) (*[]byte, error) {
	var data []byte
	var err error
	isJSON := IsJSONFile(name) // - Проверка что это json расширение файла err
	if isJSON {
		data, err = json.Marshal(binList)
		if err != nil {
			return nil, err
		}
	}
	return &data, nil
}
