package logger

type Config struct {
	AccessLogLevel       string `env:"HE_FLOORS_ACCESS_LOG_LEVEL" env-default:"INFO"`
	ErrorLogLevel        string `env:"HE_FLOORS_ERROR_LOG_LEVEL" env-default:"ERROR"`
	AccessLogDestination string `env:"HE_FLOORS_ACCESS_LOG_DESTINATION" env-default:"stdout"`
	ErrorLogDestination string `env:"HE_FLOORS_ERROR_LOG_DESTINATION" env-default:"stderr"`
}
