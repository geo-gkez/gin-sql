package errors

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ErrorHandlerMiddleware combines panic recovery and error handling
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer recoverFromPanics(c)
		c.Next()
		handleErrors(c)
	}
}

// recoverFromPanics handles any panics during request processing
func recoverFromPanics(c *gin.Context) {
	if err := recover(); err != nil {
		log.Println("Recovered from panic:", err)
		renderErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", fmt.Sprintf("%v", err))
		c.Abort()
	}
}

// handleErrors processes errors added via ctx.Error()
func handleErrors(c *gin.Context) {
	if len(c.Errors) == 0 {
		return
	}

	err := c.Errors.Last().Err
	log.Printf("Error: %v", err)

	// Check for custom error types
	var appErr AppError
	if errors.As(err, &appErr) {
		// Use the status code from our custom error
		renderErrorResponse(c, appErr.StatusCode, http.StatusText(appErr.StatusCode), appErr.Message)
		return
	}

	// Default error handling
	renderErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
}

// renderErrorResponse creates consistent error responses with complete problem details
func renderErrorResponse(c *gin.Context, status int, title, detail string, options ...func(*ProblemDetails)) {
	problemDetails := ProblemDetails{
		Status: status,
		Title:  title,
	}

	// Set detail if provided
	if detail != "" {
		problemDetails.Detail = detail
	}

	// Apply any additional options to set other fields
	for _, option := range options {
		option(&problemDetails)
	}

	// Set default type if not provided
	if problemDetails.Type == "" {
		problemDetails.Type = fmt.Sprintf("https://httpstatuses.com/%d", status)
	}

	// Set instance to current request path if not provided
	if problemDetails.Instance == "" {
		problemDetails.Instance = c.Request.URL.Path
	}

	c.JSON(status, problemDetails)
}
