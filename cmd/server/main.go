package main

import (
	"bobobox/config"
	"bobobox/internal/router"
	"bobobox/pkg/app"
	"bobobox/pkg/log"
)

func main() {
	logger := log.New()

	db, err := config.InitDB() // Initialize the database using sqlx
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	// App Init
	appConfig := app.AppConfig{
		Db:     db,
		Logger: logger,
	}

	router.Router(appConfig)
}
