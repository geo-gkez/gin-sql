package main

import (
	"fmt"
	"org/gg/banking/internal/config"
)

func main() {
	router, serverConfig := config.SetupApp()

	err := router.Run(fmt.Sprintf(":%d", serverConfig.Port))
	if err != nil {
		return
	}
}
