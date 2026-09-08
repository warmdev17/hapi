package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessEnvelope wraps a successful response.
type SuccessEnvelope struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// ErrorEnvelope wraps an error response.
type ErrorEnvelope struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Details any    `json:"details"`
}

// Success write.s a standardized success response.
func Success(c *gin.Context, statusCode int, data any, message string) {
	if message == "" {
		message = http.StatusText(statusCode)
	}

	c.JSON(statusCode, SuccessEnvelope{
		Status:  "success",
		Code:    statusCode,
		Message: message,
	})
}

// Error writes a standardized error response with a machine-readable code.
func Error(c *gin.Context, statusCode int, message string, machineCode string) {
	c.JSON(statusCode, ErrorEnvelope{
		Status:  "error",
		Code:    statusCode,
		Message: message,
		Error:   machineCode,
		Details: nil,
	})
}

// ErrorCode maps HTTP status to machine-readable error code.
func ErrorCode(status int) string {
	switch status {
	case http.StatusBadRequest:
		return "BAD_REQUEST"
	case http.StatusUnauthorized:
		return "UNAUTHORIZED"
	case http.StatusForbidden:
		return "FORBIDDEN"
	case http.StatusNotFound:
		return "NOT_FOUND"
	case http.StatusConflict:
		return "CONFLICT"
	case http.StatusRequestEntityTooLarge:
		return "REQUEST_TOO_LARGE"
	case http.StatusTooManyRequests:
		return "RATE_LIMITED"
	case http.StatusBadGateway:
		return "UPSTREAM_ERROR"
	case http.StatusGatewayTimeout:
		return "TIMEOUT"
	case http.StatusInternalServerError:
		return "INTERNAL_ERROR"
	default:
		return "UNKNOWN_ERROR"
	}
}

// ErrorSimple writes an error with machine code derived from status.
func ErrorSimple(c *gin.Context, status int, message string) {
	Error(c, status, message, ErrorCode(status))
}

// ErrorValidation writes a 400 validation error with specific field details.
func ErrorValidation(c *gin.Context, message string, field string) {
	c.JSON(http.StatusBadRequest, ErrorEnvelope{
		Status:  "error",
		Code:    http.StatusBadRequest,
		Message: message,
		Error:   "VALIDATION_ERROR",
		Details: map[string]string{"field": field},
	})
}
