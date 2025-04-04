package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"org/gg/banking/models/errors"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recovered from error:", err)
				problemDetails := errors.ProblemDetails{
					Status: http.StatusInternalServerError,
					Title:  "Internal Server Error",
				}
				c.JSON(http.StatusInternalServerError, problemDetails)
			}
		}()
		c.Next()
	}
}
