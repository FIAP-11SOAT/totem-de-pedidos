package api

import (
	"totem-pedidos/internal/api/routers"

	"github.com/labstack/echo/v4"
)

func Routers(e *echo.Echo) {
	routers.HealthRouter(e)
}
