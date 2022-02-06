package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func isClientError(code int) bool {
	return code/100 == 4
}

func isServerError(code int) bool {
	return code/100 == 5
}

// Response is a general response type used for http requests
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse is returned in case of error
func ErrorResponse(c *gin.Context, code int, message string, err error) {
	if isClientError(code) {
		zap.L().Warn(message, zap.Int("status", code), zap.Error(err))
	} else if isServerError(code) {
		zap.L().Error(message, zap.Int("status", code), zap.Error(err))
	}
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

// SuccessResponse is returned after a successful request
func SuccessResponse(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
