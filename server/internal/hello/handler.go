package hello

import (
	"fmt"
	"net/http"
)

type HalloHandler struct{}

func NewHalloHandler(router *http.ServeMux) {
	handler := &HalloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}

func (handler *HalloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")            // В консоль
		fmt.Fprint(w, "<h1>Hello</h1>") // В браузер
	}
}
