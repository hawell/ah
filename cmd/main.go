package main

import (
	"ah/server"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

func main() {
	// load configs for main package from environment variables
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		log.Fatal(err)
	}

	httpServer, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	httpServer.ListenAndServe()
}
