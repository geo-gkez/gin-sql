package main

import (
	"org/gg/banking/internal/config"
)

func main() {
	router := config.SetupApp()

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
