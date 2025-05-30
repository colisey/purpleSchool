package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	t := time.Now()
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
// 	numGoroutines := 3
// 	chSize := len(arr) / numGoroutines

// 	results := make(chan int, chSize)
// 	var wg sync.WaitGroup

// 	for i := 0; i < numGoroutines; i++ {
// 		start := i * chSize
// 		end := start + chSize

// 		if i == numGoroutines-1 {
// 			end = len(arr)
// 		}

// 		wg.Add(1)

// 		go func(chunk []int) {
// 			defer wg.Done()
// 			results <- arrSum(chunk)
// 		}(arr[start:end])
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	total := 0
// 	for sum := range results {
// 		total += sum
// 		fmt.Printf("Сумма промежуточная: %d\n", total)
// 	}
// 	fmt.Printf("Сумма массива: %d\n", total)
// 	fmt.Printf("И того: %s\n", time.Since(t))
// }

// func arrSum(arr []int) int {
// 	result := 0
// 	for _, res := range arr {
// 		result += res
// 	}
// 	return result
// }
