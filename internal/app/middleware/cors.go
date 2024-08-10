package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	DaySeconds = 86400
)

// NewCORS returns a new CORS middleware.
func NewCORSMiddleware() echo.MiddlewareFunc {
	config := middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:*",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"Cache-Control",
			"X-Requested-With",
			"X-CSRF-Token",
			"accept",
		},
		AllowMethods: []string{
			echo.GET,
			echo.HEAD,
			echo.POST,
		},
		AllowCredentials: true,
		MaxAge:           DaySeconds,
	}

	return middleware.CORSWithConfig(config)
}
