package controllers

import (
	models "crudecho/model"
	"crudecho/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkerController struct {
	service *services.WorkerService
}

func NewWorkerController(service *services.WorkerService) *WorkerController {
	return &WorkerController{service: service}
}

func (c *WorkerController) CreateWorker(ctx echo.Context) error {
	var worker models.Worker
	if err := ctx.Bind(&worker); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	result, err := c.service.CreateWorker(worker)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, result)
}

func (c *WorkerController) GetWorkerByID(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	worker, err := c.service.GetWorkerByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, worker)
}

func (c *WorkerController) GetAllWorkers(ctx echo.Context) error {
	workers, err := c.service.GetAllWorkers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, workers)
}

func (c *WorkerController) UpdateWorker(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	var update models.Worker
	if err := ctx.Bind(&update); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	result, err := c.service.UpdateWorker(id, update)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, result)
}

func (c *WorkerController) DeleteWorker(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	result, err := c.service.DeleteWorker(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, result)
}
