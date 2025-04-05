package main

import (
	"fmt"
	"org/gg/banking/internal/config"
)

func main() {
	router, serverConfig := config.SetupApp()

	//TODO user serverConfig.Mode to set the mode
	err := router.Run(fmt.Sprintf(":%d", serverConfig.Port))
	if err != nil {
		return
	}
}
