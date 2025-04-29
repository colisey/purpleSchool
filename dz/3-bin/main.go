package main

import (
	"main/storage"

	"github.com/fatih/color"
)

func main() {
	binlist, _ := storage.Read("data.json")
	if binlist != nil {
		for _, bin := range binlist.Bins {
			color.Green(bin.Name)

		}
	}
}
