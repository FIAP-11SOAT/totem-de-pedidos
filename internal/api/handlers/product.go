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

func (p *ProductHandler) ListAllProducts(c echo.Context) error {
	// implement-me
	return c.JSON(http.StatusInternalServerError, "")
}

func (p *ProductHandler) FindProductById(c echo.Context) error {
	// implement-me
	return c.JSON(http.StatusInternalServerError, "")
}

// CreateProduct is the handler to receive IO and call usecase to create a product.
func (p *ProductHandler) CreateProduct(c echo.Context) error {
	productInput := new(usecase.ProductInput)
	if err := c.Bind(productInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := productInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	createdProduct, err := p.productService.CreateProduct(productInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdProduct)
}

func (p *ProductHandler) UpdateProduct(c echo.Context) error {
	// implement-me
	return c.JSON(http.StatusInternalServerError, "")
}

func (p *ProductHandler) DeleteProduct(c echo.Context) error {
	// implement-me
	return c.JSON(http.StatusInternalServerError, "")
}

func (p *ProductHandler) ListAllCategories(c echo.Context) error {
	// implement-me
	return c.JSON(http.StatusInternalServerError, "")
}
