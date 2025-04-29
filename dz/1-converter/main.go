package main

import (
	"errors"
	"fmt"
	"math"
)

const USDtoEUR = 0.91
const USDtoRUB = 84.28

func main() {
	currencies := []string{"USD", "EUR", "RUB"}
	name := map[string][]string{
		"base":   {"базовая", "базовой", "базовую"},
		"quoted": {"котируемая", "котируемой", "котируемую"},
	}
	fmt.Println("Привет! Это конвертер валют.")
	baseCurrency, baseErr := getCurrency(&currencies, name["base"])
	if baseErr != nil {
		fmt.Println(baseErr)
		return
	}
	sum := getSumm()

	filteredCurrency := getFilteredCurrencies(&currencies, baseCurrency)
	quotedCurrency, quotedErr := getCurrency(&filteredCurrency, name["quoted"])
	if quotedErr != nil {
		fmt.Println(quotedErr)
		return
	}
	converter(baseCurrency, quotedCurrency, sum)
}
func getCurrency(currencies *[]string, print []string) (string, error) {
	var curr string
	var error error = nil
baseLoop:
	for {
		fmt.Printf("Введите %s валюту: ", print[2])
		curr = scanCurrency(currencies)
		// fmt.Printf("Вы ввели %s\n", curr)
		switch {
		case contains(currencies, curr):
			fmt.Printf("Вы выбрали %s валютой %s\n", print[1], curr)
			break baseLoop
		case curr == "Exit":
			fmt.Println("До свидания!")
			error = errors.New("Вы завершили операцию")
			break baseLoop
		default:
			fmt.Println("Неверная валюта. Попробуйте еще раз.")
			continue
		}
	}
	return curr, error
}

func getSumm() int {
	var sum int
baseLoop:
	for {
		fmt.Print("Введите сумму: ")
		fmt.Scan(&sum)
		if sum > 0 {
			fmt.Printf("Вы ввели сумму %d\n", sum)
			break baseLoop
		} else {
			fmt.Println("Сумма должна быть положительной. Попробуйте еще раз.")
			continue
		}
	}
	return sum
}
func getFilteredCurrencies(currencies *[]string, current string) []string {
	var filteredCurrencies []string

	for _, currency := range *currencies {
		if currency != current {
			filteredCurrencies = append(filteredCurrencies, currency)
		}
	}
	return filteredCurrencies
}

func converter(baseCurrName string, quotedCurrName string, sum int) (float32, error) {
	// var convertedSum float32
	var error error = nil
	type ConvertMap map[string]float32

	var EURtoRUB = math.Round((USDtoRUB/USDtoEUR)*100) / 100
	convertTo := baseCurrName + "to" + quotedCurrName

	convertMap := ConvertMap{
		"USDtoRUB": float32(sum) * USDtoRUB,
		"RUBtoUSD": float32(sum) / USDtoRUB,
		"EURtoRUB": float32(float64(sum) * EURtoRUB),
		"RUBtoEUR": float32(float64(sum) / EURtoRUB),
		"EURtoUSD": float32(sum) / USDtoEUR,
		"USDtoEUR": float32(sum) * USDtoEUR,
	}
	//TODO Как в мапу типизировать функцию? map{onSumm: sum => float32(sum) * USDtoRUB }

	// switch convertTo {
	// case "USDtoRUB":
	// 	convertedSum = float32(sum) * USDtoRUB
	// case "RUBtoUSD":
	// 	convertedSum = float32(sum) / USDtoRUB
	// case "EURtoRUB":
	// 	convertedSum = float32(float64(sum) * EURtoRUB)
	// case "RUBtoEUR":
	// 	convertedSum = float32(float64(sum) / EURtoRUB)
	// case "EURtoUSD":
	// 	convertedSum = float32(sum) / USDtoEUR
	// case "USDtoEUR":
	// 	convertedSum = float32(sum) * USDtoEUR
	// default:
	// 	error = errors.New("Неизвестная базовая валюта")
	// }

	if error != nil {
		return 0, error
	}

	fmt.Printf("Конвертированная сумма в валюте %s: %.2f\n", convertTo, convertMap[convertTo])
	return convertMap[convertTo], nil
}

func scanCurrency(currencies *[]string) string {
	var curr string
	for index, currency := range *currencies {
		separator := ", "
		if (len(*currencies) - 1) <= index {
			separator = "\n"
		}
		fmt.Printf("%s%s", currency, separator)
	}
	fmt.Scan(&curr)
	return curr
}
func contains(s *[]string, e string) bool {
	for _, a := range *s {
		if a == e {
			return true
		}
	}
	return false
}
