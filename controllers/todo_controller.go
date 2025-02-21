package controllers

import (
	"net/http"

	"example/hello/config"
	"example/hello/models"

	"github.com/labstack/echo/v4"
)

type TodoController struct{}

func NewTodoController() *TodoController {
	return &TodoController{}
}

func (uc *TodoController) GetTodos(c echo.Context) error {
	var todos []models.Todo
	config.DB.Find(&todos)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    todos,
	})
}

func (uc *TodoController) GetTodo(c echo.Context) error {
	id := c.Param("id")
	var todo models.Todo
	config.DB.First(&todo, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    todo,
	})
}

func (uc *TodoController) CreateTodo(c echo.Context) error {
	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if err := config.DB.Create(&todo); err.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed",
			"error":   err.Error,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    todo,
	})
}
