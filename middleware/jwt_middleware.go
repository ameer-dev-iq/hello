package middleware

// import (
// 	"github.com/golang-jwt/jwt"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// func JWTMiddleware() echo.MiddlewareFunc {
// 	// Define the JWT configuration
// 	config := middleware.JWTConfig{
// 		SigningKey: []byte("secret"), // Replace with your secret key
// 		// Parse the token into jwt.MapClaims
// 		Claims: jwt.MapClaims{},
// 		// Custom error handler for invalid tokens
// 		ErrorHandlerWithContext: func(err error, c echo.Context) error {
// 			return echo.NewHTTPError(echo.ErrUnauthorized.Code, "Invalid or expired token")
// 		},
// 	}

// 	// Return the JWT middleware
// 	return middleware.JWTWithConfig(config)
// }
