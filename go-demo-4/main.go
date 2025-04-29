package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault(files.NewJsonDb("data.json"))
	// vault := account.NewVault(cloud.NewCloudDB("https://a.ru"))
Menu:
	for {
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func findAccount(vault *account.VaultWithDb) {

	url := promptData([]string{"Введите url"})
	accounts := vault.FindAccountsByURL(url)

	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
		return
	}
	for _, account := range accounts {
		account.Output()
	}
}
func deleteAccount(vault *account.VaultWithDb) {
	counts := len(vault.Accounts)
	url := promptData([]string{"Введите url"})
	total := vault.DeleteAccountByUrl(url)
	if total < counts {
		color.Red("Аккаунтов наудалял %d штуки", counts-total)
	}
}
func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		if err.Error() == "INVALID_URL" {
			output.PrintError("Неверный формат URL")
		} else if err.Error() == "INVALID_LOGIN" {
			output.PrintError("Неверный формат LOGIN")
		}
		return
	}

	vault.AddAccount(*myAccount)
}

func promptData[T any](prompts []T) string {

	for index, prompt := range prompts {
		if index == len(prompts)-1 {
			fmt.Printf("%v: ", prompt) // %v Форматирует любой тип
		} else {
			fmt.Println(prompt)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
