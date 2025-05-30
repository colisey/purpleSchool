package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ping(url string, respChan chan int, errChan chan error) {
	resp, err := http.Get(url)
	if err != nil {
		errChan <- err
		return
		// fmt.Printf("Ошибка запроса %s", err.Error())
	}
	respChan <- resp.StatusCode
}

func main() {
	path := flag.String("file", "url.txt", "path to URL file")
	flag.Parse()
	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}
	urlSlice := strings.Split(string(file), "\r\n")
	respChan := make(chan int)
	errChan := make(chan error)
	for _, url := range urlSlice {
		go ping(url, respChan, errChan)
	}
	for range urlSlice {
		select {
		case errRes := <-errChan:
			fmt.Println(errRes)
		case resp := <-respChan:
			fmt.Println(resp)
		}
	}
}
