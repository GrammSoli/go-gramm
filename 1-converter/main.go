package main

import "fmt"

const (
	usdToEur = 0.88
	usdToRub = 79.13
	eurToRub = usdToRub / usdToEur
)

func main() {

	fmt.Println("1 EUR в RUB")

	fmt.Print(eurToRub)
}
