package output

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(value any) {
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибка: %d", t)
	case error:
		color.Red("Ошибка: %v", t)
	default:
		color.Red("Неизвестная ошибка: %v", t)
	}
}

func sum[T int | string](a, b T) T {
	switch d := any(a).(type) {
	case string:
		fmt.Println("Type of a:", d)
	}
	return a + b
}

type List[T any] struct {
	Items []T
}

func (l *List[T]) AddElement(item T) {
	l.Items = append(l.Items, item)
}
