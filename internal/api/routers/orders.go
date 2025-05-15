package routers

import (
	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/labstack/echo/v4"
)

func OrdersRouter(e *echo.Echo, dbConnection *dbadapter.DatabaseAdapter) {}
