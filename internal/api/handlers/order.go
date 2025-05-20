package handlers

import (
	"net/http"
	"strconv"

	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/input"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/mapper"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/usecase"
	"github.com/labstack/echo/v4"
)

type OrderHanlder struct {
	orderService usecase.Order
}

func NewOrderHandler(service usecase.Order) *OrderHanlder {
	return &OrderHanlder{
		orderService: service,
	}
}

func (o *OrderHanlder) CreateOrder(c echo.Context) error {
	var orderInput input.OrderInput
	if err := c.Bind(&orderInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if err := orderInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	orderEntity := mapper.ToOrderDomain(orderInput)

	orderId, err := o.orderService.CreateOrder(orderEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create order"})
	}

	return c.JSON(
		http.StatusCreated,
		map[string]string{"message": "order created successfully", "orderId": strconv.Itoa(orderId)},
	)
}

func (o *OrderHanlder) UpdateOrder(c echo.Context) error {
	var input input.UpdateOrderInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	if err := o.orderService.UpdateOrderStatus(id, input.Status); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update order"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "order updated successfully"})
}

func (o *OrderHanlder) GetOrderById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid order ID"})
	}

	order, err := o.orderService.GetOrderByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "order not found"})
	}

	return c.JSON(http.StatusOK, order)
}

func (o *OrderHanlder) GetOrders(c echo.Context) error {
	var filter input.OrderFilterInput
	if err := c.Bind(&filter); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid query parameters"})
	}

	orders, err := o.orderService.ListOrders(filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch orders"})
	}

	return c.JSON(http.StatusOK, orders)
}

func (o *OrderHanlder) Checkout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid order ID"})
	}

	if err := o.orderService.Checkout(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "payment failed or could not update order"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "payment successful"})
}
