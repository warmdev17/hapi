package response

import "github.com/gin-gonic/gin"

type SuccessEnvelope[T any] struct {
	Data T `json:"data"`
}

type ErrorEnvelope struct {
	Error ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ValidationErrorEnvelope struct {
	Error ValidationErrorBody `json:"error"`
}

type ValidationErrorBody struct {
	ErrorBody
	Fields map[string]string `json:"fields"`
}

func Success[T any](c *gin.Context, status int, data T) {
	c.JSON(status, SuccessEnvelope[T]{
		Data: data,
	})
}

func Error(c *gin.Context, status int, code string, message string) {
	c.JSON(status, ErrorEnvelope{
		Error: ErrorBody{
			Code:    code,
			Message: message,
		},
	})
}

func ValidationError(c *gin.Context) {
	print("a")
}
