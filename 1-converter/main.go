package main

import "fmt"

func main() {
	fmt.Println("1 EUR в RUB")
	usdToEur, usdToRub := getUserInput()
	eurToRub := (usdToRub / usdToEur)
	fmt.Print(eurToRub)
}

func getUserInput() (float64, float64) {
	var usdToEur float64
	var usdToRub float64
	fmt.Print("Введите курс USD в EUR: ")
	fmt.Scan(&usdToEur)
	fmt.Print("Ведите курс USD в RUB: ")
	fmt.Scan(&usdToRub)
	return usdToEur, usdToRub
}

func currency(number int, usd float64, eur float64) float64 {

}
