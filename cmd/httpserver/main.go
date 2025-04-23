package main

import (
	"totem-pedidos/internal/adapters/postgres"
	"totem-pedidos/internal/api"
	domain "totem-pedidos/internal/core/domain/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// chama o adapter do postgres
	postgres := postgres.New()

	// domain order
	orderDomain := domain.New()

	orderHandler := api.NewOrderRouter(postgres, orderDomain)
	orderHandler.OrderRouters(e)

	e.Logger.Fatal(e.Start(":1323"))
}
