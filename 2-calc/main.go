package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	promt_choise  = "AVG - среднее, SUM - сумма, MED - медиана"
	promt_numbers = "Ведите числа в  строку, через запятую (1,2)"
)

func main() {
	// Проверим, что аргументы командной строки переданы
	if len(os.Args) < 3 {
		fmt.Println("Неверное количество аргументов. Пример использования: ./script <операция> <числа через запятую>")
		return
	}

	// Получаем операцию и числа из аргументов
	choice := strings.ToUpper(os.Args[1])
	numbers := os.Args[2]

	// Проверяем правильность выбранной операции
	if choice != "AVG" && choice != "SUM" && choice != "MED" {
		fmt.Println("Неверный выбор операции. Используйте AVG, SUM или MED.")
		return
	}

	// Рассчитываем результат
	result, err := calculate(choice, numbers)
	if err != nil {
		fmt.Println("Ошибка при расчете:", err)
		return
	}

	// Выводим результат
	fmt.Printf("Ваш(-а) %s: %.0f\n", choice, result)
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
