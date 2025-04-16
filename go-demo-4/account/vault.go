package account

import (
	"demo/password/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
	Total     int       `json:"total"`
}

func (vault *Vault) DeleteAccountByUrl(url string) int {
	var accounts []Account

	// isDelited := false
	for _, account := range vault.Accounts {

		if !strings.Contains(account.Url, url) {
			accounts = append(accounts, account)
			continue
		}
		fmt.Println("Удалил аккаунты " + account.Login)
		// isDelited = true
	}
	vault.Accounts = accounts
	vault.save()
	return len(accounts)
}
func (vault *Vault) FindAccountsByURL(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {

		if strings.Contains(account.Url, url) {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()
	vault.Total = len(vault.Accounts)
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}
