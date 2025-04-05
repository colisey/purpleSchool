package main

import "fmt"

const USDtoEUR = 0.91
const USDtoRUB = 84.28

func main() {
	sum := getSum()
	fmt.Printf("Вы ввели %.2f EUR, это будет %.2f RUB\n", sum, eurToRub(sum))
}
func eurToRub(sum float32) float32 {
	return sum / USDtoEUR * USDtoRUB
}

func getSum() float32 {
	fmt.Println("Введите сумму в EUR:")
	var sum float32
	fmt.Scanln(&sum)
	return sum
}
