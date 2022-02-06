package server

import (
	"ah/logger"
	"ah/server/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func newRouter(accessLogger *zap.Logger, storage interface{}) *gin.Engine {
	handleRecovery := func(c *gin.Context, err interface{}) {
		handlers.ErrorResponse(c, http.StatusInternalServerError, err.(string), nil)
		c.Abort()
	}
	router := gin.New()
	router.Use(gin.CustomRecovery(handleRecovery))
	router.Use(logger.MiddlewareFunc(accessLogger))
	router.NoMethod(func(ctx *gin.Context) {
		handlers.ErrorResponse(ctx, http.StatusMethodNotAllowed, "requested method is not allowed", nil)
	})
	router.NoRoute(func(ctx *gin.Context) {
		handlers.ErrorResponse(ctx, http.StatusNotFound, "path not found", nil)
	})
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	})
	router.Use(func(ctx *gin.Context){
		ctx.Set("db", storage)
	})
	router.GET("get_providers", handlers.GetProviders)
	return router
}
