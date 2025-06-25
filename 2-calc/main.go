package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Проверим, что аргументы командной строки переданы
	if len(os.Args) < 3 {
		fmt.Println("Неверное количество аргументов. Пример использования: ./calc <операция> <числа через запятую>")
		return
	}

	// Получаем операцию и числа из аргументов
	choice := strings.ToUpper(os.Args[1])
	numbers := os.Args[2]

	// Разбираем строку чисел
	intNumbers, err := parseNumbers(numbers)
	if err != nil {
		fmt.Println("Ошибка при преобразовании:", err)
		return
	}

	// Получаем функцию расчета
	opFunc, ok := operations[choice]
	if !ok {
		fmt.Println("Неверный выбор операции. Используйте AVG, SUM или MED.")
		return
	}

	// Рассчитываем результат
	result, err := opFunc(intNumbers)
	if err != nil {
		fmt.Println("Ошибка при расчете:", err)
		return
	}

	// Выводим результат
	fmt.Printf("Ваш(-а) %s: %.0f\n", choice, result)
}

// parseNumbers преобразует строку чисел вида "1,2,3" в срез int
func parseNumbers(numbers string) ([]int, error) {
	if numbers == "" {
		return nil, errors.New("введены пустые числа")
	}

	parts := strings.Split(numbers, ",")
	intNumbers := make([]int, 0, len(parts))
	for _, numStr := range parts {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		intNumbers = append(intNumbers, num)
	}
	return intNumbers, nil
}

// operations содержит список доступных операций
var operations = map[string]func([]int) (float64, error){
	"AVG": func(nums []int) (float64, error) {
		if len(nums) == 0 {
			return 0, errors.New("нет чисел для расчета среднего")
		}
		sum := 0
		for _, n := range nums {
			sum += n
		}
		return float64(sum) / float64(len(nums)), nil
	},
	"SUM": func(nums []int) (float64, error) {
		sum := 0
		for _, n := range nums {
			sum += n
		}
		return float64(sum), nil
	},
	"MED": func(nums []int) (float64, error) {
		if len(nums) == 0 {
			return 0, errors.New("нет чисел для расчета медианы")
		}
		sort.Ints(nums)
		n := len(nums)
		if n%2 == 1 {
			return float64(nums[n/2]), nil
		}
		return float64(nums[n/2-1]+nums[n/2]) / 2, nil
	},
}
