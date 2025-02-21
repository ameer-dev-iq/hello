package main

import (
	"example/hello/config"
	"example/hello/controllers"
	"example/hello/helper"
	"log"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize Echo
	e := echo.New()

	// Custom Validator
	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	if v, ok := e.Validator.(*helper.CustomValidator); ok {
		v.Validator.RegisterValidation("custom", func(fl validator.FieldLevel) bool {
			return true
		})
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize Database
	config.InitDB()

	// Routes
	authController := controllers.NewAuthController()
	userController := controllers.NewUserController()
	todoController := controllers.NewTodoController()

	// Public routes (no authentication required)
	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)

	// Restricted routes (require JWT authentication)
	r := e.Group("/users")
	// Apply JWT middleware
	r.GET("", userController.GetUsers)
	r.GET("/:id", userController.GetUser)

	// apply jwt
	todo := e.Group("/todos")

	todo.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))

	todo.GET("", todoController.GetTodos)
	todo.GET("/:id", todoController.GetTodo)
	todo.POST("", todoController.CreateTodo)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
