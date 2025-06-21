package storage

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"go-gramm/3-password/account"
	"go-gramm/3-password/files"
)

// Bin описывает данные BIN.
type Bin struct {
	Accounts  *[]account.Account `json:"accounts"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

// SaveBin сохраняет BIN в файл JSON по указанному пути.
func SaveBin(path string, bin Bin) error {
	if !files.IsJSONFile(path) {
		return errors.New("wrong file type")
	}
	bins, _ := ReadBins(path)
	bins = append(bins, bin)
	data, err := json.MarshalIndent(bins, "", "    ")
	if err != nil {
		return err
	}
	files.WriteFile(data, path)
	return nil
}

// ReadBins читает список BIN из файла JSON.
func ReadBins(path string) ([]Bin, error) {
	if !files.IsJSONFile(path) {
		return nil, errors.New("wrong file type")
	}
	data, err := files.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Bin{}, nil
		}
		return nil, err
	}
	if len(data) == 0 {
		return []Bin{}, nil
	}
	var bins []Bin
	err = json.Unmarshal(data, &bins)
	if err != nil {
		return nil, err
	}
	return bins, nil
}
