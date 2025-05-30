package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// func main() {
// 	code := make(chan int)
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			getHttpCode(code)
// 			wg.Done()
// 		}()
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(code)
// 	}()
// 	for res := range code {
// 		fmt.Printf("Status code = %d\n", res)
// 	}
// 	// res := <-code
// }

// func getHttpCode(codeCn chan int) {
// 	resp, err := http.Get("https://google.com")
// 	if err != nil {
// 		fmt.Printf("Ошибка: %s\n", err.Error())
// 		return
// 	}
// 	defer resp.Body.Close()

// 	codeCn <- resp.StatusCode
// }
