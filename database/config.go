package database

// Config holds database specific configurations
type Config struct {
	ConnectionString string `env:"AH_FLOORS_DATABASE_CONNECTION_STRING" env-default:"flooruser:floorpass@tcp(127.0.0.1:3306)/floor"`
}
