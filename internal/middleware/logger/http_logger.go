package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"org/gg/banking/internal/config/logger"
	"strings"

	"github.com/gin-gonic/gin"
)

// ResponseBodyWriter is a custom ResponseWriter that captures the response body
type ResponseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write captures the response and writes it to the original writer
func (r *ResponseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// HTTPLoggerMiddleware logs both incoming requests and outgoing responses
func HTTPLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before request
		logIncomingRequest(c)

		// Create a custom ResponseWriter to capture the response
		rbw := &ResponseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = rbw

		// Process the request
		c.Next()

		// After request - log the response
		logOutgoingResponse(c, rbw)
	}
}

// logIncomingRequest logs the details of an incoming HTTP request
func logIncomingRequest(c *gin.Context) {
	// Save the request body so it can be read multiple times
	bodyValue := captureRequestBody(c)

	logger.Logger.InfoContext(
		c,
		"INCOMING REQUEST",
		slog.String("method", c.Request.Method),
		slog.String("path", c.Request.URL.Path),
		slog.String("query", c.Request.URL.RawQuery),
		slog.Any("headers", c.Request.Header),
		slog.Any("body", bodyValue), // Use slog.Any to handle JSON objects properly
	)
}

// captureRequestBody reads the request body and restores it for subsequent handlers
func captureRequestBody(c *gin.Context) any {
	var bodyBytes []byte
	if c.Request.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(c.Request.Body)

		if err != nil {
			logger.Logger.ErrorContext(c, "Error reading request body", slog.Any("error", err))
			return nil
		}
		// Restore the request body for subsequent handlers
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	return parseBody(bodyBytes, c.GetHeader("Content-Type"))
}

// parseBody converts raw bytes to structured data based on content type
func parseBody(bodyBytes []byte, contentType string) any {
	// Handle empty bodies
	if len(bodyBytes) == 0 {
		return ""
	}

	// If Content-Type is JSON, parse it as JSON
	if strings.Contains(contentType, "application/json") {
		jsonData, success := parseJSON(bodyBytes)
		if success {
			return jsonData
		}
	}

	// Default case: return as string
	return string(bodyBytes)
}

// parseJSON attempts to parse bytes as JSON, returning success flag
func parseJSON(data []byte) (any, bool) {
	var jsonData any
	if err := json.Unmarshal(data, &jsonData); err == nil {
		return jsonData, true
	}
	return string(data), false
}

// truncateString cuts a string if it exceeds maxLength
func truncateString(s string, maxLength int) string {
	if len(s) > maxLength {
		return s[:maxLength] + "... (truncated)"
	}
	return s
}

// logOutgoingResponse logs the details of the HTTP response
func logOutgoingResponse(c *gin.Context, rbw *ResponseBodyWriter) {
	// Get status code and size
	statusCode := rbw.Status()
	responseSize := rbw.body.Len()

	// Get response headers
	responseHeaders := rbw.Header()

	// Process the response body
	contentType := responseHeaders.Get("Content-Type")
	var bodyValue any

	// Only process non-empty bodies
	if responseSize > 0 {
		responseBody := rbw.body.Bytes()
		bodyValue = processResponseBody(responseBody, contentType)
	} else {
		bodyValue = "<empty body>"
	}

	// Log the response
	logger.Logger.InfoContext(
		c,
		"OUTGOING RESPONSE",
		slog.Int("status", statusCode),
		slog.Int("size", responseSize),
		slog.Any("headers", responseHeaders),
		slog.Any("body", bodyValue),
		slog.String("path", c.Request.URL.Path),
		slog.String("method", c.Request.Method),
	)
}

// processResponseBody handles formatting response body based on content type
func processResponseBody(responseBody []byte, contentType string) any {
	// If content type is JSON, parse it as JSON
	if strings.Contains(contentType, "application/json") {
		jsonData, success := parseJSON(responseBody)
		if success {
			return jsonData
		}
	}

	// For non-JSON content types, convert to string and truncate if needed
	bodyStr := string(responseBody)
	return truncateString(bodyStr, 1000)
}
