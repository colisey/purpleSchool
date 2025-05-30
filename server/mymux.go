package main

// import (
// 	"fmt"
// 	"net/http"
// )

// // MyMux — собственный роутер.
// type MyMux struct {
// 	routes map[string]http.Handler
// }

// // ServeHTTP реализует интерфейс http.Handler.
// func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	handler, exists := m.routes[r.URL.Path]
// 	if exists {
// 		handler.ServeHTTP(w, r)
// 		return
// 	}
// 	http.NotFound(w, r)
// }

// // Метод для регистрации маршрутов.
// func (m *MyMux) Handle(path string, handler http.Handler) {
// 	m.routes[path] = handler
// }

// // Вспомогательная функция для регистрации http.HandlerFunc
// func (m *MyMux) HandleFunc(path string, handlerFunc http.HandlerFunc) {
// 	m.routes[path] = handlerFunc
// }

// // Пример использования.
// func main() {
// 	mux := &MyMux{routes: make(map[string]http.Handler)}

// 	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "Hello from /hello")
// 	})

// 	mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "Goodbye from /bye")
// 	})

// 	fmt.Println("Server is running on :8080")
// 	http.ListenAndServe(":8080", mux)
// }
