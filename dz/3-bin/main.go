package main

import (
	"main/storage"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		color.Red("Error loading .env file")
	}

	binlist, _ := storage.Read("data.json")
	if binlist != nil {
		for _, bin := range binlist.Bins {
			color.Green(bin.Name)

		}
	}
}
