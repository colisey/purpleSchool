package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"

	"github.com/fatih/color"
)

// Zoo
// Animal -> Crocodile

func main() {
	// fmt.Println(rand.IntN(10))

	login := promptdata("Введите логин")
	password := promptdata("Введите пароль")
	url := promptdata("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)

	if err != nil {
		if err.Error() == "INVALID_URL" {
			color.Red("Неверный формат URL")
		} else if err.Error() == "INVALID_LOGIN" {
			color.Magenta("Неверный формат LOGIN")
		}
		return
	}

	// fmt.Println(myAccount.generatePassword(12))
	// myAccount.generatePassword(12)
	myAccount.OutputPassword()

	files.ReadFile()
	files.WriteFile()

	fmt.Println(myAccount)
}

func promptdata(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
