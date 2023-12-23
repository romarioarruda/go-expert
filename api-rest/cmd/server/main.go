package main

import (
	"fmt"
	"go-expert/api-rest/configs"
)

func main() {
	config, _ := configs.LoadConfig()

	fmt.Printf("Config %v", config)
}
