package ui

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func (h *handler) GetHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
