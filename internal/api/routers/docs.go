package routers

import (
	"totem-pedidos/internal/handlers"

	"github.com/labstack/echo/v4"
)

func DocsRouter(e *echo.Echo) {
	h := handlers.NewDocsHandler()
	e.GET("/docs/:path", h.Docs)
}
