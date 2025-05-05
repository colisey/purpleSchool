package geo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}
	// resp, err := http.Get("https://ipapi.co/json")
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, err
	}
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
