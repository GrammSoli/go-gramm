package files

import (
	"fmt"
	"go-gramm/3-password/output"
	"os"
	"path/filepath"
	"strings"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(filename string) *JsonDb {
	return &JsonDb{
		filename: filename,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError("Ошибка создания файла: " + err.Error())
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.PrintError("Ошибка записи в файл: " + err.Error())
		return
	}

	fmt.Println("Запись прошла успешно!")
}

func (db *JsonDb) IsJSONFile(name string) bool {
	return strings.EqualFold(filepath.Ext(db.filename), ".json")
}
