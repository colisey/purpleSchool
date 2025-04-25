package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateProductID() string {
	currentTime := time.Now().UnixNano()
	randomNum := rand.Intn(1000) // случайное число
	return fmt.Sprintf("%d-%d", currentTime, randomNum)
}
