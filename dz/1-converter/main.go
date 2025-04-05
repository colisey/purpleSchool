package main

import "fmt"

const USDtoEUR = 0.91
const USDtoRUB = 84.28

func main() {
	fmt.Println(eurToRub(10))
}
func eurToRub(sum float32) float32 {
	return sum / USDtoEUR * USDtoRUB
}

// Для начала работы с калькулятором, нам понадобится:
// - Создать новый проект
// - Объявить пакет main
// - Объявить функцию main
// - Создать константы конвертации
// - Из USD в EUR
// - Из USD в RUB
// - Рассчитать EUR в RUB на основании первых двух
// В git создать ветку: 1-start
// Папка: 1-converter
