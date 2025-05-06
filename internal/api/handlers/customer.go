package handlers

import (
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
		return c.JSON(http.StatusBadRequest, "error binding create customer request")
	}

	h.CustomerService.CreateCustomer(&customerCreateRequest)
	return c.JSON(http.StatusOK, "")
}
