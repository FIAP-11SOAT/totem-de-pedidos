package routers

import (
	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api/handlers"
	"github.com/labstack/echo/v4"
)

func CustomerRouter(e *echo.Echo, dbConnection *dbadapter.DatabaseAdapter) {
	h := handlers.NewCustomerHandler(dbConnection)

	e.POST("/customer", h.CreateCustomer)
	e.GET("/customer", h.IdentifyCustomer)
}
