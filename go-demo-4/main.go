package main

import (
	"fmt"
)

func main() {
	// a := 5
	// double(&a)
	// fmt.Println(a)
	a := [4]int{1, 2, 3, 4}
	reverse(&a)
	fmt.Println(a)
}

// func double(num *int) {
// 	*num = *num * 2
// }

func reverse(arr *[4]int) {
	// arr2 := *arr
	// length := len(*arr) - 1
	// for i := 0; i <= length; i++ {
	// 	arr[i] = arr2[length-i]
	// }
	for index, value := range *arr {
		(*arr)[len(arr)-1-index] = value
		// arr[len(arr)-1-index] = value
	}

}
