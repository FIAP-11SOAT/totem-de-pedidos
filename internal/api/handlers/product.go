package handlers

import (
	"net/http"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	service "github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/usecase"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/usecase"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/repositories"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService usecase.Product
}

func NewProductHandler(dbConnection *dbadapter.DatabaseAdapter) *ProductHandler {
	productRepository := repositories.NewProductRepository(dbConnection)

	return &ProductHandler{
		productService: service.NewProductUseCase(productRepository),
	}
}

func (h *ProductHandler) ListAllProducts(c echo.Context) error {
	// implement-me
	return c.JSON(http.StatusInternalServerError, "")

}

func (h *ProductHandler) FindProductById(c echo.Context) error {
	// implement-me
	return c.JSON(http.StatusInternalServerError, "")

}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	// implement-me
	return c.JSON(http.StatusInternalServerError, "")

}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	// implement-me
	return c.JSON(http.StatusInternalServerError, "")
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	// implement-me
	return c.JSON(http.StatusInternalServerError, "")
}

func (h *ProductHandler) ListAllCategories(c echo.Context) error {
	// implement-me
	return c.JSON(http.StatusInternalServerError, "")
}
