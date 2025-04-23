package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateOrder(c echo.Context) error {

	// instancia um domain

	return c.String(http.StatusOK, "Hello, World!")
}
