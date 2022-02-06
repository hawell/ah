package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates two new loggers, one for access log and one for error log
func NewLogger() (*zap.Logger, *zap.Logger, error) {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		return nil, nil, err
	}

	accessLogLevel := zapcore.InfoLevel
	_ = accessLogLevel.UnmarshalText([]byte(config.AccessLogLevel))

	zapAccessLogConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(accessLogLevel),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "",
			MessageKey:     "message",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{config.AccessLogDestination},
		ErrorOutputPaths: []string{config.ErrorLogDestination},
	}
	accessLogger, err := zapAccessLogConfig.Build()
	if err != nil {
		return nil, nil, err
	}

	errorLogLevel := zapcore.InfoLevel
	_ = errorLogLevel.UnmarshalText([]byte(config.AccessLogLevel))

	zapErrorLogConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(errorLogLevel),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{config.ErrorLogDestination},
		ErrorOutputPaths: []string{config.ErrorLogDestination},
	}
	errorLogger, err := zapErrorLogConfig.Build()
	if err != nil {
		return nil, nil, err
	}

	return accessLogger, errorLogger, nil
}

// MiddlewareFunc logs every request to access log
func MiddlewareFunc(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Info("",
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.EscapedPath()),
		)
		ctx.Next()
	}
}
