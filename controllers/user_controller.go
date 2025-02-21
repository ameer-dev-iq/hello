package controllers

import (
	"example/hello/config"
	"example/hello/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUsers(c echo.Context) error {
	var users []models.User
	if err := c.Bind(&users); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	config.DB.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func (uc *UserController) GetUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, user)
}
