package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponce struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recover ", r)
		}
	}()
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			panic("Такого города нет: " + city)
		}
		return &GeoData{
			City: city,
		}, nil
	}
	// resp, err := http.Get("https://ipapi.co/json")
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println(resp)
		return nil, errors.New("NOT200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, err
	}
	return &geo, nil
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var populationResponce CityPopulationResponce
	json.Unmarshal(body, &populationResponce)
	return !populationResponce.Error
}

// go func handler(w http.ResponseWriter, r *http.Request) {
// 	tags := r.URL.Query()["tag"] // Все значения параметра "tag" fmt.Fprintf(w, "%v", tags)
// 	// // Например, ["one" "two"]

// 	for _, tag := range tags {
// 		fmt.Println(tag)
// 	}
// 	}
