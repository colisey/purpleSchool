package bins

import (
	"encoding/json"
	"time"
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

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		Id:       id,
		Private:  private,
		CreateAt: time.Now(),
		Name:     name,
	}
}
