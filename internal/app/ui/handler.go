package ui

import (
	echo "github.com/labstack/echo/v4"
)

type Handler interface {
	GetHealth(c echo.Context) error
}

type handler struct{}

func NewHandler() Handler {
	return &handler{}
}
