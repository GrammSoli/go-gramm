package account

import (
	"encoding/json"
	"go-gramm/3-password/output"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write(content []byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"createdAt"`
}
type VaultWithDb struct {
	Vault
	db Db
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
		output.PrintError("Ошибка десериализации данных: " + err.Error())
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

func (vault *VaultWithDb) FindAccountsByUrl(url string) []Account {
	url = strings.TrimSpace(strings.ToLower(url))
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	str = strings.TrimSpace(strings.ToLower(str))
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := checker(account, str)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted
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
	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.PrintError("Ошибка сериализации данных: " + err.Error())
	}
	vault.db.Write(data)
}
