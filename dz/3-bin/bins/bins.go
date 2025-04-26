package bins

import (
	"encoding/json"
	"main/utils"
	"time"

	"github.com/fatih/color"
)

// В bins структура и методы реализованы корректно, но стоит улучшить именование полей
// и добавить больше экспортируемых функций для взаимодействия с другими пакетами.

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

func (binlist *BinList) AddBinList(bin Bin) {
	binlist.Bins = append(binlist.Bins, bin)
	binlist.UpdatedAt = time.Now()
}

func (binlist *BinList) ToBytes(name string) (*[]byte, error) {
	// Сначала надо преобразовать в file в зависимости от типа
	// file, err := json.Marshal(binlist)
	file, err := utils.ConvertToBytes(&binlist, name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		Id:       id,
		Private:  private,
		CreateAt: time.Now(),
		Name:     name,
	}
}
func NewBinList(data []byte) *BinList {
	var binList BinList
	err := json.Unmarshal(data, &binList)

	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
	}
	return &binList
}
