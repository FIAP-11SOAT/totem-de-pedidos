package routers

import (
	"github.com/labstack/echo/v4"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api/handlers"
)

func PaymentsRouter(e *echo.Echo, dbConnection *dbadapter.DatabaseAdapter) {
	h := handlers.PaymentHandler{}
	// r := dbadapter.NewPaymentsRepository(dbConnection)
	// u := usecase.NewPaymentsUseCase(r)

	g := e.Group("/payments")
	g.GET("/:id", h.GetByID)
	g.POST("/webhook", h.PaymentWebHook)

}
