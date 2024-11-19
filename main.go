package main

import (
	"crudecho/config"
	"crudecho/controllers"
	models "crudecho/model" // Import the models package
	"crudecho/routes"
	"crudecho/services"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Manager{}); err != nil {
		log.Fatal("Migration failed: ", err)
	}
}

func main() {
	// Load the environment variables and connect to databases
	cfg := config.Init()

	// Run migrations
	migrate(cfg.PostgresDB)

	// Initialize Echo instance
	e := echo.New()

	// Initialize services
	workerService := services.NewWorkerService(cfg)
	managerService := services.NewManagerService(cfg)

	// Initialize controllers with the services
	workerController := controllers.NewWorkerController(workerService)
	managerController := controllers.NewManagerController(managerService)

	// Set up routes
	routes.SetupRoutes(e, workerController, managerController)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
