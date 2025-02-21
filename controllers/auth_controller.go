package controllers

import (
	"example/hello/config"
	"example/hello/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (controller AuthController) Login(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	var dbUser models.User
	if err := config.DB.Where("username = ?", user.Username).First(&dbUser); err == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid username",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid username or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       dbUser.ID,
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
		"data":    t,
	})
}

func (controller AuthController) Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{Message: err.Error()})
	}

	var dbUser models.User
	if err := config.DB.Where("username = ?", user.Username).First(&dbUser); err.Error == nil {
		return c.JSON(http.StatusUnauthorized,
			models.BaseResponse[string]{Message: "Username already exists"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user.Password = string(hashedPassword)

	config.DB.Create(&user)
	return c.JSON(http.StatusOK, models.BaseResponse[string]{Message: "success", Data: user.Username})
}
