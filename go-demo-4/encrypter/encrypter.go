package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// type Encrypter interface {
// 	Encrypt(string) string
// 	Decrypt(string) string
// }

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	defer func() { // Обработка паники, выведет текст 33 строки например
		r := recover()
		if r != nil {
			fmt.Println("Recover ", r)
		}
	}()
	key := os.Getenv("KEY")
	// os.Setenv("MY_VAR", "value")
	if key == "" {
		panic("Не передан параметр KEY в переменные окружения")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encriptt(plainStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key)) // Создаём новый блок для фифрования
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block) // Передаём его для создания нового GSM
	if err != nil {
		panic(err.Error())
	}
	// nonce =  number once used. Здесь хранится уникальное значение для безопастности криптграфичесских операций
	nonce := make([]byte, aesGSM.NonceSize()) // Формируем некоторый случайный нонс, который мы используем для шифрования
	_, err = io.ReadFull(rand.Reader, nonce)  // С помощью rand.Reader
	if err != nil {
		panic(err.Error())
	}
	return aesGSM.Seal(nonce, nonce, plainStr, nil) // Seal шифрует. Запечатываем шифром

}
func (enc *Encrypter) Decriptt(encryptedStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key)) // Создаём новый блок для фифрования
	if err != nil {
		color.Red("block, err := aes.NewCipher")
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block) // Передаём его для создания нового GSM
	if err != nil {
		color.Red("aesGSM, err := cipher.NewGCM(block)")
		panic(err.Error())
	}
	nonceSize := aesGSM.NonceSize()
	nonce, cipherText := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	plainText, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		color.Red("plainText, err := aesGSM.Open(nil, nonce, cipherText, nil)")
		panic(err.Error())
	}
	return plainText

}
