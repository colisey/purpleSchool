package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по login",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("___Менеджер паролей___")
	// res := os.Getenv("VAR")
	// fmt.Println(res)

	// for _, e := range os.Environ() { // Вывод переменных env среды
	// 	fmt.Println(e)
	// }

	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())
	// vault := account.NewVault(cloud.NewCloudDB("https://a.ru"))
Menu:
	for {
		variant := promptData(menuVariants...)

		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }
	}
}

func findAccountByUrl(vault *account.VaultWithDb) {

	url := promptData("Введите url")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}
func findAccountByLogin(vault *account.VaultWithDb) {

	login := promptData("Введите login")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
		return
	}
	for _, account := range *accounts {
		account.Output()
	}

}

//	func checkUrl(acc account.Account, str string) bool {
//		return strings.Contains(acc.Url, str)
//	}
func deleteAccount(vault *account.VaultWithDb) {
	counts := len(vault.Accounts)
	url := promptData("Введите url")
	total := vault.DeleteAccountByUrl(url)
	if total < counts {
		color.Red("Аккаунтов наудалял %d штуки", counts-total)
	}
}
func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

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

func promptData(prompts ...string) string {

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
