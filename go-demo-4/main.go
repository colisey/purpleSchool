package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

// Zoo
// Animal -> Crocodile

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createAt time.Time
	updateAt time.Time
	account
}

func (acc account) outputPassword() {
	// fmt.Println(acc)
	// fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

// func newAccount(login, password, urlString string) (*account, error) { // 9.10. Функция constructor
//
//		// Валидируем аккаунт
//		if login == "" {
//			return nil, errors.New("INVALID_LOGIN")
//		}
//		_, err := url.ParseRequestURI(urlString)
//		if err != nil {
//			return nil, errors.New("INVALID_URL")
//		}
//		newAcc := &account{
//			login:    login,
//			password: password,
//			url:      urlString,
//		}
//		if password == "" {
//			newAcc.generatePassword(12)
//		}
//		return newAcc, nil
//	}
func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) { // 9.10. Функция constructor
	// Валидируем аккаунт
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &accountWithTimeStamp{
		createAt: time.Now(),
		updateAt: time.Now(),
		account: account{
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

var letterRunes = []rune("qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM1234567890-*!")

func main() {
	// fmt.Println(rand.IntN(10))

	login := promptdata("Введите логин")
	password := promptdata("Введите пароль")
	url := promptdata("Введите URL")

	myAccount, err := newAccountWithTimeStamp(login, password, url)

	if err != nil {
		if err.Error() == "INVALID_URL" {
			fmt.Println("Неверный формат URL")
		} else if err.Error() == "INVALID_LOGIN" {
			fmt.Println("Неверный формат LOGIN")
		}
		return
	}

	// fmt.Println(myAccount.generatePassword(12))
	// myAccount.generatePassword(12)
	myAccount.outputPassword()
	fmt.Println(myAccount)
}

func promptdata(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

// type auto struct {
// 	brand string
// 	model string
// 	year  int
// }
// func (a auto) printAuto() {
// 	fmt.Printf("%s %s, %d", a.brand, a.model, a.year)
// }
// func (a *auto) updateYear(n int) {
// 	a.year = n
// }

// func sdsd()  {
// 	toyota := auto{
// 		"Toyota"
// 		"Camry"
// 		2020
// 	}

// 	toyota.updateYear(2021)
// }

// user := struct {name string age int}{ name: "Fofof", age: 21}

// Вопрос: Как создать структуру Engine с полем power типа int и затем встроить её в структуру Car,
// имеющую также поле brand типа string? Сможешь привести пример кода?

type Engine struct {
	power int
}
type Car struct {
	brand string
	dors  int
	Engine
}

func (a *Car) addDors(n int) {
	a.dors = n
}

func (a *Car) upPower(n int) {
	a.power += n
}

func (c Car) printDoors() {
	fmt.Printf("Количество дверей: %d\n", c.dors)
}
