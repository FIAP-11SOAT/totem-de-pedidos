package routers

import (
	"github.com/labstack/echo/v4"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api/handlers"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/usecase"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/repositories"
)

func OrdersRouter(e *echo.Echo, dbConnection *dbadapter.DatabaseAdapter) {
	or := repositories.NewOrderRepository(dbConnection)
	pr := repositories.NewProductRepository(dbConnection)
	u := usecase.NewOrderUseCase(or, pr)
	h := handlers.NewOrderHandler(u)

	e.GET("/orders", h.GetOrders)
	e.GET("/orders/:id", h.GetOrderById)
	e.POST("/orders", h.CreateOrder)
	e.PUT("/orders/:id", h.UpdateOrder)
}
