package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var typeOperations string
	var numbersOperations string
	for {
		fmt.Println("Введите тип операции:\n AVG - среднее,\n SUM - сумму,\n MED - медиану)\n")
		fmt.Scan(&typeOperations)
		fmt.Println("typeOperations == ", typeOperations == "SUM")
		if typeOperations == "AVG" || typeOperations == "SUM" || typeOperations == "MED" {
			break
		} else {
			fmt.Println("Вы ввели неправельный тип операции, введите заново")
		}
	}

	fmt.Println("Введите числа через запятую с которыми будем работать\n")
	fmt.Scan(&numbersOperations)
	convertedArgs := convertStringToArgs(numbersOperations)
	fmt.Println(calculate(typeOperations, convertedArgs...))
	// fmt.Println(calculate("AVG", 2, 10, 9))
	// fmt.Println(calculate("SUM", 2, 10, 9))
	// fmt.Println(calculate("MED", 2, 10, 9, 7, 1, 3, 4))
}

func calculate(operation string, args ...int) int {
	switch operation {
	case "AVG":
		return sum(args...) / len(args)
	case "SUM":
		return sum(args...)
	case "MED":
		sortedArgs := sortArray(args)
		midIndex := len(sortedArgs) / 2
		if len(sortedArgs)%2 == 0 {
			return (sortedArgs[midIndex-1] + sortedArgs[midIndex]) / 2
		} else {
			return sortedArgs[midIndex]
		}
	default:
		return 0
	}
}
func sum(args ...int) int {
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}

func sortArray(args []int) []int {
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args)-i-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}
	return args

}

func convertStringToArgs(args string) []int {
	strArgs := strings.Split(args, ",")
	intArgs := make([]int, 0, len(strArgs))
	for _, value := range strArgs {
		n, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			fmt.Println("Ошибка преобразования:", err)
			break
		}
		intArgs = append(intArgs, n)
	}
	return intArgs
}

// Array и Slice
// Нужно создать приложение, которое
// - Принимает операцию (AVG - среднее, SUM - сумму, MED -
// медиану)
// - Принимает неограниченное число чисел через запятую (2, 10,
// 9)
// - Разбивает строку чисел по запятым и затем делает расчёт в
// зависимости от операции выводя результат
