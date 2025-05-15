package routers

import (
	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/labstack/echo/v4"
)

func CategoryRouter(e *echo.Echo, dbConnection *dbadapter.DatabaseAdapter) {
	e.GET("/categories", func(c echo.Context) error {
		return c.JSON(200, "ok")
	})
}
