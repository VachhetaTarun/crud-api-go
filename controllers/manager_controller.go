package controllers

import (
	models "crudecho/model"
	"crudecho/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ManagerController struct {
	service *services.ManagerService
}

func NewManagerController(service *services.ManagerService) *ManagerController {
	return &ManagerController{service: service}
}

func (c *ManagerController) CreateManager(ctx echo.Context) error {
	var manager models.Manager
	if err := ctx.Bind(&manager); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	err := c.service.CreateManager(manager)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, manager)
}

func (c *ManagerController) GetManagerByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	manager, err := c.service.GetManagerByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, manager)
}

func (c *ManagerController) GetAllManagers(ctx echo.Context) error {
	managers, err := c.service.GetAllManagers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, managers)
}

func (c *ManagerController) UpdateManager(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	var update models.Manager
	if err := ctx.Bind(&update); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	err = c.service.UpdateManager(uint(id), update)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, update)
}

func (c *ManagerController) DeleteManager(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	err = c.service.DeleteManager(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"result": "success"})
}
