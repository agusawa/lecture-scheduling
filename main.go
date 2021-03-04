package main

import (
	"lecture-scheduling/application"
	"lecture-scheduling/config"
	"lecture-scheduling/repository"
	"lecture-scheduling/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Parse flags to env
	config.InitFlags()

	// Setup configuration
	configuration := config.New("./.env")
	database := config.NewSqliteDatabase(configuration)
	defer database.Close()

	// Setup repository
	scheduleRepository := repository.NewScheduleRepository(database)

	// Setup service
	scheduleService := service.NewScheduleService(&scheduleRepository)

	// Setup and run the app
	app := application.New(&scheduleService)
	app.Run()
}
