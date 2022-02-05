package main

// Config is used to get configs from environment variables
type Config struct {
	LogLevel    string `env:"HE_FLOORS_LOG_LEVEL" env-default:"INFO"`
}

