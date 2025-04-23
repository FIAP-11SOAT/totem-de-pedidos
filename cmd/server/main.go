package main

import (
	"totem-pedidos/internal/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	api.Routers(e)
	e.Logger.Fatal(e.Start(":5000"))
}
