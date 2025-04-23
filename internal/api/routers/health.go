package routers

import (
	"totem-pedidos/internal/handlers"

	"github.com/labstack/echo/v4"
)

func HealthRouter(e *echo.Echo) {
	h := handlers.NewHealthHandler()
	e.GET("/health", h.HealthCheck)
}
