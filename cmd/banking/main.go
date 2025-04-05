package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"org/gg/banking/internal/config"
)

func main() {
	router, serverConfig := config.SetupApp()

	gin.SetMode(serverConfig.Mode)
	err := router.Run(fmt.Sprintf(":%d", serverConfig.Port))
	if err != nil {
		return
	}
}
