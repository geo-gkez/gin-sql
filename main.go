package main

import (
	"org/gg/banking/config"
)

func main() {
	router := config.SetupApp()

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
