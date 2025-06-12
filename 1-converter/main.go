package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	usdToEur     = 0.86
	usdToRub     = 79.0
	promt_from   = "Введите исходную валюту (USD/EUR/RUB):"
	promt_amount = "Введите сумму: "
	promt_to     = "Введите целевую валюту (USD/EUR/RUB):"
)

func main() {
	for {
		fmt.Println("Валютный калькулятор")
		from, err := getUserInputCurrency(promt_from)
		if err != nil {
			fmt.Println("Неверная валюта, попробуйте еще раз.")
			continue
		}

		amount, err := getUserInputAmount(promt_amount)
		if err != nil {
			fmt.Println("Неверная сумма, попробуйте еще раз.")
			continue
		}
		available := getAvailableCurrencies(from)
		promt_to := fmt.Sprintf("Введите целевую валюту (%s):", available)
		to, err := getUserInputCurrency(promt_to)
		if err != nil || to == from {
			fmt.Println("Неверная валюта, попробуйте еще раз.")
			continue
		}
		fmt.Println("Вы выбрали исходную валюту:", from)
		fmt.Println("Ваша сумма:", amount)
		fmt.Println("Вы выбрали целевую валюту:", to)
		result := convert(from, to, amount)
		fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)
		isRepeateCalculation := checkRepeatCallculation()
		if !isRepeateCalculation {
			break
		}
	}
}

func getUserInputCurrency(prompt string) (string, error) {
	fmt.Println(prompt)
	var curr string
	fmt.Scan(&curr)
	curr = strings.ToUpper(curr)
	if curr != "USD" && curr != "EUR" && curr != "RUB" {
		return "", errors.New("NO_PARAMETRS")
	}
	return curr, nil
}

func getUserInputAmount(prompt string) (float64, error) {
	fmt.Println(prompt)
	var amou float64
	fmt.Scan(&amou)
	if amou <= 0 {
		return 0, errors.New("NO_PARAMETRS")
	}
	return amou, nil
}

func getAvailableCurrencies(from string) string {
	switch from {
	case "USD":
		return "EUR/RUB"
	case "EUR":
		return "USD/RUB"
	case "RUB":
		return "USD/EUR"
	default:
		return "USD/EUR/RUB"
	}
}

func convert(from string, to string, amount float64) float64 {
	eurToRub := usdToRub / usdToEur

	if from == to {
		return amount
	}
	switch from {
	case "USD":
		switch to {
		case "EUR":
			return amount * usdToEur
		case "RUB":
			return amount * usdToRub
		}
	case "EUR":
		switch to {
		case "USD":
			return amount / usdToEur
		case "RUB":
			return amount * eurToRub
		}
	case "RUB":
		switch to {
		case "USD":
			return amount / usdToRub
		case "EUR":
			return amount / eurToRub
		}
	}
	return 0
}

func checkRepeatCallculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще рассчет? (Y/n): ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
