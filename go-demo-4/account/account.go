package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM1234567890-*!")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createAt time.Time
	updateAt time.Time
	Account
}

func (acc Account) OutputPassword() {
	// fmt.Println(acc)
	// fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}
func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) { // 9.10. Функция constructor
	// Валидируем аккаунт
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &AccountWithTimeStamp{
		createAt: time.Now(),
		updateAt: time.Now(),
		Account: Account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
