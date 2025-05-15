package handlers

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/inputs"
	"net/http"
	"strconv"

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
	filter := new(inputs.ProductFilterInput)
	if err := c.Bind(filter); err != nil {
		c.Logger().Error("Error binding filter", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	products, err := p.productService.GetProducts(filter)
	if err != nil {
		c.Logger().Error("Error getting products", err)
		return err
	}

	if len(products) == 0 {
		return c.JSON(http.StatusNotFound, make([]string, 0))
	}

	return c.JSON(http.StatusOK, products)
}

func (p *ProductHandler) FindProductById(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, "")
}

func (p *ProductHandler) CreateProduct(c echo.Context) error {
	productInput := new(inputs.ProductInput)
	if err := c.Bind(productInput); err != nil {
		c.Logger().Error("Error binding product input", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := productInput.Validate(); err != nil {
		c.Logger().Error("Error binding product input", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	createdProduct, err := p.productService.CreateProduct(productInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdProduct)
}

func (p *ProductHandler) UpdateProduct(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, "")
}

func (p *ProductHandler) DeleteProduct(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, "")
}

func (p *ProductHandler) ListAllCategories(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, "")
}

func (p *ProductHandler) GetProductByCategoryID(c echo.Context) error {
	CategoryID := c.Param("id")
	if CategoryID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Category ID is required"})
	}

	CategoryIDInt, err := strconv.Atoi(CategoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Category ID"})
	}

	products, erro := p.productService.GetProductByCategoryID(CategoryIDInt)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": erro.Error()})
	}

	return c.JSON(http.StatusOK, products)
}
