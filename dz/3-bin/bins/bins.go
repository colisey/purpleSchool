package bins

import (
	"encoding/json"
	"errors"
	"main/files"
	"time"

	"github.com/fatih/color"
)

type Bin struct {
	Id       string    `json:"ID"`
	Private  bool      `json:"private"`
	CreateAt time.Time `json:"createAt"`
	Name     string    `json:"name"`
}

type BinList struct { // список BinList
	Bins      []Bin     `json:"bins"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (bl *BinList) AddBinList(bin Bin) {
	bl.Bins = append(bl.Bins, bin)
}

func (binlist *BinList) ToBytes() ([]byte, error) {
	file, err := json.Marshal(binlist)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func (binlist *BinList) save(name string) error {
	isJSON := files.IsJSONFile(name) // - Проверка что это json расширение файла err
	if !isJSON {
		color.Red("Файл %s не является файлом *.json", name)
		return errors.New("DONT_JSON")
	}

	data, err := binlist.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data, name)
	return nil
}

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		Id:       id,
		Private:  private,
		CreateAt: time.Now(),
		Name:     name,
	}
}
