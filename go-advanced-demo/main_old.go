package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// func main() {
// 	const requestCount = 10
// 	var wg sync.WaitGroup
// 	t := time.Now()

// 	for i := 0; i < requestCount; i++ {
// 		wg.Add(1)
// 		go func() {
// 			getHttpCode(i, t)
// 			wg.Done()
// 		}()
// 	}
// 	fmt.Printf("Цикл завершился: %s\n", time.Since(t))
// 	// time.Sleep(time.Second * 3)
// 	wg.Wait()
// 	fmt.Printf("И того: %s\n", time.Since(t))
// }

// func getHttpCode(i int, t time.Time) {
// 	resp, err := http.Get("https://google.com")
// 	if err != nil {
// 		fmt.Printf("Request %d Ошибка: %s\n", i, err.Error())
// 		return
// 	}
// 	defer resp.Body.Close()

// 	fmt.Printf("Request %d: status code = %d, Время = %s \n", i, resp.StatusCode, time.Since(t))
// }
