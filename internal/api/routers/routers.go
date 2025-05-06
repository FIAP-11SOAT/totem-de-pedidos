package routes

import (
	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/api/handlers"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	echoEngine *echo.Echo
}

func New(echoEngine *echo.Echo) *Routes {
	return &Routes{
		echoEngine: echoEngine,
	}
}

func (r *Routes) SetupRouters(dbConnection *dbadapter.DatabaseAdapter) *echo.Echo {
	defaultHandler := handlers.NewHealthHandler()
	productHandler := handlers.NewProductHandler(dbConnection)
	customerHandler := handlers.NewCustomerHandler(dbConnection)

	r.setupHandlers(
		defaultHandler,
		productHandler,
		customerHandler,
	)

	return r.echoEngine
}

func (r *Routes) setupHandlers(
	defaultHandler *handlers.HealthHandler,
	productHandler *handlers.ProductHandler,
	customerHandler *handlers.CustomerHandler,
) {
	r.echoEngine.GET("/health", defaultHandler.HealthCheck)

	// product endpoints
	r.echoEngine.GET("/categories", productHandler.ListAllCategories)
	r.echoEngine.GET("/products", productHandler.ListAllProducts)
	r.echoEngine.GET("/products/:id", productHandler.FindProductById)
	r.echoEngine.POST("/products", productHandler.CreateProduct)
	r.echoEngine.PUT("/products/:id", productHandler.UpdateProduct)
	r.echoEngine.DELETE("/products/:id", productHandler.DeleteProduct)
	r.echoEngine.POST("/customer", customerHandler.CreateCustomer)
}
