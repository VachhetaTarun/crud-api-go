package routes

import (
	"crudecho/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, workerController *controllers.WorkerController, managerController *controllers.ManagerController) {
	// Worker routes
	e.POST("/workers", workerController.CreateWorker)
	e.GET("/workers/:id", workerController.GetWorkerByID)
	e.GET("/workers", workerController.GetAllWorkers) // Add this line
	e.PUT("/workers/:id", workerController.UpdateWorker)
	e.DELETE("/workers/:id", workerController.DeleteWorker)

	// Manager routes
	e.POST("/managers", managerController.CreateManager)
	e.GET("/managers/:id", managerController.GetManagerByID)
	e.GET("/managers", managerController.GetAllManagers) // Add this line
	e.PUT("/managers/:id", managerController.UpdateManager)
	e.DELETE("/managers/:id", managerController.DeleteManager)
}
