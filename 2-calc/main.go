package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	promt_choise  = "AVG - среднее, SUM - сумма, MED - медиана"
	promt_numbers = "Ведите числа в  строку, через запятую (1,2)"
)

func main() {
	for {
		choice, err := getUserchoosingOperation(promt_choise)
		if err != nil {
			fmt.Println("Неверный выбор, попробуйте еще раз.")
			continue
		}
		numbers := getUserParseNumbers(promt_numbers)
		result, err := calculate(choice, numbers)
		fmt.Printf("Ваш(-а) %s: %.0f", choice, result)
		isRepeateCalculation := checkRepeatCallculation()
		if !isRepeateCalculation {
			break
		}
	}
}

func getUserchoosingOperation(promt_choise string) (string, error) {
	fmt.Println(promt_choise)
	var choice string
	fmt.Scan(&choice)
	choice = strings.ToUpper(choice)
	if choice != "AVG" && choice != "SUM" && choice != "MED" {
		return "", errors.New("NO_PARAMETRS")
	}
	return choice, nil
}

func getUserParseNumbers(promt_numbers string) string {
	fmt.Println(promt_numbers)
	var numbers string
	fmt.Scan(&numbers)
	return numbers
}

func calculate(choice string, numbers string) (float64, error) {
	str := strings.Split(numbers, ",")

	var intNumbers []int
	for _, numStr := range str {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Ошибка при преобразовании:", err)
			return 0, err
		}
		intNumbers = append(intNumbers, num)
	}

	switch choice {
	case "AVG":

		var sum int
		for _, num := range intNumbers {
			sum += num
		}
		avg := float64(sum) / float64(len(intNumbers))
		return avg, nil

	case "SUM":

		var sum int
		for _, num := range intNumbers {
			sum += num
		}
		return float64(sum), nil

	case "MED":

		sort.Ints(intNumbers)
		n := len(intNumbers)
		var median float64
		if n%2 == 1 {
			median = float64(intNumbers[n/2])
		} else {
			median = float64(intNumbers[n/2-1]+intNumbers[n/2]) / 2
		}
		return median, nil
	}
	return 0, nil
}
func checkRepeatCallculation() bool {
	var userChoise string
	fmt.Println("Вы хотите сделать еще рассчет? (Y/n): ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
