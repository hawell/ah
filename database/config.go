package database

type Config struct {
	ConnectionString string `env:"AH_DATABASE_CONNECTION_STRING" env-default:"root:root@tcp(127.0.0.1:3306)/floor"`
}