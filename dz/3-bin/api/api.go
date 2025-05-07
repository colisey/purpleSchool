package api

import (
	"fmt"
	"main/config"
)

func api() {
	config := config.NewConfig()
	fmt.Println(config)
}
