package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("New project")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")
	// format := flag.Bool("format", 1, "Формат вывода погоды")
	flag.Parse()

	fmt.Println(*city)
	// fmt.Println(*format)

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)

	// r := strings.NewReader("Привет, я поток данных")
	// block := make([]byte, 4)
	// for {
	// 	_, err := r.Read(block)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Printf("%q\n", block)
	// }
}
