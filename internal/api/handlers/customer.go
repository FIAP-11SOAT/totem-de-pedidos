package handlers

import (
	"fmt"
	"net/http"

	dbadapter "github.com/FIAP-11SOAT/totem-de-pedidos/internal/adapters/database"
	service "github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/usecase"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/usecase"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/repositories"
	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	CustomerService usecase.Customer
}

func NewCustomerHandler(dbConnection *dbadapter.DatabaseAdapter) *CustomerHandler {
	CustomerRepository := repositories.NewCustomerRepository(dbConnection)

	return &CustomerHandler{
		CustomerService: service.NewCustomerUseCase(CustomerRepository),
	}
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	var customerCreateRequest usecase.CustomerInput
	if err := c.Bind(&customerCreateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error binding create customer request"})
	}

	customerResponse, err := h.CustomerService.CreateCustomer(&customerCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, customerResponse)
}

func (h *CustomerHandler) IdentifyCustomer(c echo.Context) error {
	taxId := c.QueryParam("taxid")
	fmt.Println(taxId)

	if taxId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "param taxid missing"})
	}

	customerResponse, err := h.CustomerService.IdentifyCustomer(&taxId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if customerResponse == nil {
		fmt.Println("no content")
		return c.JSON(http.StatusNoContent, "")
	}

	return c.JSON(http.StatusOK, customerResponse)
}
