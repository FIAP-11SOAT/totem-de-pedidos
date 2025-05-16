package handlers

import (
	"net/http"
	"strconv"

	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/input"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/mapper"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/output"

	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/usecase"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService usecase.Product
}

func NewProductHandler(productService usecase.Product) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (p *ProductHandler) ListAllProducts(c echo.Context) error {
	filter := new(input.ProductFilterInput)
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

	resultProducts := make([]output.ProductOutput, 0)
	for _, product := range products {
		resultProducts = append(resultProducts, mapper.MapProductToOutput(product))
	}

	return c.JSON(http.StatusOK, resultProducts)
}

func (p *ProductHandler) FindProductById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product ID is required"})
	}

	product, err := p.productService.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	return c.JSON(http.StatusOK, product)
}

func (p *ProductHandler) CreateProduct(c echo.Context) error {
	productInput := new(input.ProductInput)
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
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product ID is required"})
	}

	productInput := new(input.ProductInput)
	if err := c.Bind(productInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := productInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updatedProduct, err := p.productService.UpdateProduct(id, productInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedProduct)
}

func (p *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product ID is required"})
	}

	err := p.productService.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
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
