package routers

import (
	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api/handlers"
	"github.com/labstack/echo/v4"
)

func ProductsRouter(e *echo.Echo, dbConnection *dbadapter.DatabaseAdapter) {
	h := handlers.NewProductHandler(dbConnection)

	e.GET("/categories", h.ListAllCategories)
	e.GET("/products", h.ListAllProducts)
	e.GET("/products/:id", h.FindProductById)
	e.POST("/products", h.CreateProduct)
	e.PUT("/products/:id", h.UpdateProduct)
	e.DELETE("/products/:id", h.DeleteProduct)
	e.GET("/products/category/:id", h.GetProductByCategoryID)
}
