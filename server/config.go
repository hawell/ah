package server

// Config contains api server configurations
type Config struct {
	ListenAddress string `env:"AH_FLOORS_HTTP_LISTEN_ADDRESS" env-default:"localhost:8000"`
	ReadTimeout   uint   `env:"AH_FLOORS_SERVER_READ_TIMEOUT" env-default:"5"`
	WriteTimeout  uint   `env:"AH_FLOORS_SERVER_WRITE_TIMEOUT" env-default:"5"`
}
