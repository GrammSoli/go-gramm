package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	promt_from   = "Введите исходную валюту (USD/EUR/RUB):"
	promt_amount = "Введите сумму: "
)

type conversion = map[string]map[string]float64

var rates = conversion{
	"USD": {
		"EUR": 0.86,
		"RUB": 79.0,
	},
	"EUR": {
		"USD": 1 / 0.86,
		"RUB": 79.0 / 0.86,
	},
	"RUB": {
		"USD": 1 / 79.0,
		"EUR": 0.86 / 79.0,
	},
}

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
		result := convert(from, to, amount, &rates)
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
	if toMap, ok := rates[from]; ok {
		keys := make([]string, 0, len(toMap))
		for k := range toMap {
			keys = append(keys, k)
		}
		return strings.Join(keys, "/")
	}
	return "USD/EUR/RUB"
}

func convert(from string, to string, amount float64, rates *conversion) float64 {
	if toMap, ok := (*rates)[from]; ok {
		if rate, ok := toMap[to]; ok {
			return amount * rate
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
