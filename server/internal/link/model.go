package link

import (
	"math/rand"

	"gorm.io/gorm"
)

// Как выглядит базовая структура модели пользователя (User),
// соответствующая таблице users в базе данных?
// Достаточно полей ID (uint), Name (string), Email (string).
// Напиши Go-структуру с использованием тега gorm для ID.

// type User struct {
// 	gorm.Model
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RandStringRunes(6)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
