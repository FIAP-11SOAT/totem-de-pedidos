package handlers

import "github.com/labstack/echo/v4"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c echo.Context) error {
	return c.String(200, "OK")
}
