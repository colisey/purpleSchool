package main

import (
	"demo/password/account"
	"fmt"

	"github.com/fatih/color"
)

// Zoo
// Animal -> Crocodile

// type BookmarkMap map[string]string

func main() {
	vault := account.NewVault()
	// 1. Создать аккаунт
	// 2. Найти аккаунт
	// 3. Удалть аккаунт
	// 4. Выти

	// Карта закладок
	// bookmarks := make(BookmarkMap, 5)
	// Определяем команды и их описания
Menu:
	for {
		variant := GetMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func GetMenu() int {
	var variant int
	// commands := BookmarkMap{
	// 	"1": "Создать аккаунт",
	// 	"2": "Найти аккаунт",
	// 	"3": "Удалить аккаунт",
	// 	"4": "Выход",
	// }
	fmt.Println("Введите команду:")
	fmt.Println("1 Создать аккаунт")
	fmt.Println("2 Найти аккаунт")
	fmt.Println("3 Удалить аккаунт")
	fmt.Println("4 Выход")
	// for key, value := range commands {
	// 	fmt.Println(key, ": ", value)
	// }
	fmt.Scanln(&variant)
	return variant
}
func findAccount(vault *account.Vault) {

	url := promptdata("Введите url")
	accounts := vault.FindAccountsByURL(url)

	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
		return
	}
	for _, account := range accounts {
		account.Output()
	}
}
func deleteAccount(vault *account.Vault) {
	counts := len(vault.Accounts)
	url := promptdata("Введите url")
	total := vault.DeleteAccountByUrl(url)
	if total < counts {
		color.Red("Аккаунтов наудалял %d штуки", counts-total)
	}
}
func createAccount(vault *account.Vault) {
	login := promptdata("Введите логин")
	password := promptdata("Введите пароль")
	url := promptdata("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		if err.Error() == "INVALID_URL" {
			color.Red("Неверный формат URL")
		} else if err.Error() == "INVALID_LOGIN" {
			color.Magenta("Неверный формат LOGIN")
		}
		return
	}
	// file, err := myAccount.ToBytes()

	vault.AddAccount(*myAccount)
}

func promptdata(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
