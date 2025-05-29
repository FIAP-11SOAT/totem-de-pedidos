package routers

import (
	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api/handlers"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/usecase"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/repositories"
	"github.com/labstack/echo/v4"
)

func CustomerRouter(e *echo.Echo, dbConnection *dbadapter.DatabaseAdapter) {
	customerRepository := repositories.NewCustomerRepository(dbConnection)
	customerUseCase := usecase.NewCustomerUseCase(customerRepository)
	customerHandler := handlers.NewCustomerHandler(customerUseCase)

	e.POST("/customer", customerHandler.CreateCustomer)
	e.GET("/customer", customerHandler.IdentifyCustomer)
}
