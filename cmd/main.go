package main

import (
	"ah/database"
	"ah/logger"
	"ah/server"
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
	"log"
)

func main() {
	// load configs for main package from environment variables
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		log.Fatal(err)
	}

	accessLogger, errorLogger, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	zap.ReplaceGlobals(errorLogger)

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	httpServer, err := server.NewServer(accessLogger, db)
	if err != nil {
		log.Fatal(err)
	}

	httpServer.ListenAndServe()
}
