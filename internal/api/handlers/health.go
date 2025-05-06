package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

// NewHealthHandler cria e retorna uma nova instância de HealthHandler.
func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

func (h *HealthHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
