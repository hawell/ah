package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func IsClientError(code int) bool {
	return code/100 == 4
}

func IsServerError(code int) bool {
	return code/100 == 5
}


type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrorResponse(c *gin.Context, code int, message string, err error) {
	if IsClientError(code) {
		zap.L().Warn(message, zap.Int("status", code), zap.Error(err))
	} else if IsServerError(code) {
		zap.L().Error(message, zap.Int("status", code), zap.Error(err))
	}
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

