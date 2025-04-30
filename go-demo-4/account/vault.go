package account

import (
	"demo/password/output"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}
type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
	Total     int       `json:"total"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) int {
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
func (vault *VaultWithDb) FindAccounts(url string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {

		// if strings.Contains(account.Url, url) {
		if checker(account, url) {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDb) AddAccount(acc Account) {
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

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	vault.Total = len(vault.Accounts)
	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.PrintError("Не удалось преобразовать")
	}
	vault.db.Write(data)
}

func NewVault(db Db) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		output.PrintError("Не удалось разобрать файл data.json")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}
