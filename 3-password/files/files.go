package files

import (
	"fmt"
	"os"
)

func ReadFile() {

}

func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println("Запись прошла успешно!")
}
