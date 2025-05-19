package routers

import (
	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api/handlers"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/usecase"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/services/payment"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/repositories"
	"github.com/labstack/echo/v4"
)

func OrdersRouter(e *echo.Echo, dbConnection *dbadapter.DatabaseAdapter) {
	or := repositories.NewOrderRepository(dbConnection)
	pr := repositories.NewProductRepository(dbConnection)
	ps := payment.New()
	u := usecase.NewOrderUseCase(or, pr, ps)
	h := handlers.NewOrderHandler(u)

	e.POST("/order", h.CreateOrder)
	e.GET("/order/:id", h.GetOrderById)
	e.PUT("/order/:id", h.UpdateOrder)
	e.GET("/orders/", h.GetOrders)
	e.POST("/order/:id", h.Checkout)

	// TODO: implementar rota de webhook
}
