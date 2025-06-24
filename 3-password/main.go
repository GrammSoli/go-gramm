package main

import (
	"fmt"
	"strings"

	"go-gramm/3-password/account"
	"go-gramm/3-password/files"
	"go-gramm/3-password/output"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

func menuCounter() func() {
	i := 0
	return func() {
		i++
		fmt.Println("Вызов функции номер:", i)
	}
}

func main() {
	vault := account.NewVault(files.NewJsonDb("data.json"))
	counter := menuCounter()
	// vault := account.NewVault(cloud.NewCloudDb("https://example.com/db")) - для облачной базы данных
Menu:
	for {
		counter()
		variant := promptData(
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		)
		menuFanc := menu[variant]
		if menuFanc == nil {
			break Menu
		}
		menuFanc(vault)
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

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите url: ")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Ошибка создания аккаунта: " + err.Error())
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt ...any) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска: ")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(strings.ToLower(acc.Url), strings.ToLower(str)) ||
			strings.Contains(strings.ToLower(acc.Login), strings.ToLower(str))
	})
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для удаления: "})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Аккаунт не найден")
	}
}
