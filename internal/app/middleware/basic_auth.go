package middleware

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewBasicAuth returns a new Basic Auth middleware.
func NewBasicAuthMiddleware(authUsername, authPassword string) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte(authUsername)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(authPassword)) == 1 {
			return true, nil
		}

		return false, echo.ErrUnauthorized
	})
}
