package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/usecase"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/input"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/services/payment"
	paymentmock "github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/services/payment/mock"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/repositories/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	t.Run("should create order with valid items", func(t *testing.T) {
		orderMock := mock.NewOrderRepositoryMock()
		productMock := mock.NewProductRepositoryMock()

		productMock.FindProductByIDFunc = func(ctx context.Context, id string) (*entity.Product, error) {
			return &entity.Product{
				ID:    1,
				Name:  "X-Burger",
				Price: 25.0,
			}, nil
		}

		orderMock.CreateOrderFunc = func(order entity.Order) (int, error) {
			return 123, nil
		}

		paymentMock := paymentmock.NewPaymentServiceMock()

		orderUc := usecase.NewOrderUseCase(orderMock, productMock, paymentMock)

		order := entity.Order{
			CustomerID: 1,
			Items: []entity.OrderItem{
				{ProductID: 1, Quantity: 2},
			},
		}

		orderID, err := orderUc.CreateOrder(order)

		assert.NoError(t, err)
		assert.Equal(t, 123, orderID)
	})

	t.Run("should return error if no items in order", func(t *testing.T) {
		orderUc := usecase.NewOrderUseCase(
			mock.NewOrderRepositoryMock(),
			mock.NewProductRepositoryMock(),
			paymentmock.NewPaymentServiceMock(),
		)

		order := entity.Order{
			CustomerID: 1,
			Items:      []entity.OrderItem{},
		}

		orderId, err := orderUc.CreateOrder(order)

		assert.Error(t, err)
		assert.Equal(t, -1, orderId)
	})

	t.Run("should return error if product not found", func(t *testing.T) {
		orderMock := mock.NewOrderRepositoryMock()
		productMock := mock.NewProductRepositoryMock()

		productMock.FindProductByIDFunc = func(ctx context.Context, id string) (*entity.Product, error) {
			return nil, errors.New("product not found")
		}

		paymentMock := paymentmock.NewPaymentServiceMock()

		orderUc := usecase.NewOrderUseCase(orderMock, productMock, paymentMock)

		order := entity.Order{
			CustomerID: 1,
			Items: []entity.OrderItem{
				{ProductID: 99, Quantity: 1},
			},
		}

		orderID, err := orderUc.CreateOrder(order)

		assert.Error(t, err)
		assert.Equal(t, -1, orderID)
	})
}

func TestUpdateOrderStatus(t *testing.T) {
	t.Run("should update order status successfully", func(t *testing.T) {
		orderMock := mock.NewOrderRepositoryMock()
		productMock := mock.NewProductRepositoryMock()

		orderMock.UpdateStatusFunc = func(id int, status string) error {
			return nil
		}

		paymentMock := paymentmock.NewPaymentServiceMock()

		uc := usecase.NewOrderUseCase(orderMock, productMock, paymentMock)

		err := uc.UpdateOrderStatus(1, "COMPLETED")
		assert.NoError(t, err)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		orderMock := mock.NewOrderRepositoryMock()
		productMock := mock.NewProductRepositoryMock()

		orderMock.UpdateStatusFunc = func(id int, status string) error {
			return errors.New("order not found")
		}

		paymentMock := paymentmock.NewPaymentServiceMock()

		uc := usecase.NewOrderUseCase(orderMock, productMock, paymentMock)

		err := uc.UpdateOrderStatus(99999, "CANCELLED")
		assert.Error(t, err)
		assert.EqualError(t, err, "order not found")
	})
}

func TestGetOrderByID(t *testing.T) {
	t.Run("should return order when found", func(t *testing.T) {
		expected := entity.Order{
			ID:          1,
			CustomerID:  10,
			Status:      "PENDING",
			TotalAmount: 42.00,
			Items: []entity.OrderItem{
				{ProductID: 1, Quantity: 2, Price: 21.00},
			},
		}

		orderMock := mock.NewOrderRepositoryMock()
		orderMock.GetOrderByIDFunc = func(id int) (entity.Order, error) {
			return expected, nil
		}

		productMock := mock.NewProductRepositoryMock()
		paymentMock := paymentmock.NewPaymentServiceMock()

		uc := usecase.NewOrderUseCase(orderMock, productMock, paymentMock)

		order, err := uc.GetOrderByID(1)
		assert.NoError(t, err)
		assert.Equal(t, expected, order)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		orderMock := mock.NewOrderRepositoryMock()
		orderMock.GetOrderByIDFunc = func(id int) (entity.Order, error) {
			return entity.Order{}, errors.New("not found")
		}

		productMock := mock.NewProductRepositoryMock()
		paymentMock := paymentmock.NewPaymentServiceMock()

		uc := usecase.NewOrderUseCase(orderMock, productMock, paymentMock)

		order, err := uc.GetOrderByID(999)
		assert.Error(t, err)
		assert.Equal(t, entity.Order{}, order)
	})
}

func TestListOrders(t *testing.T) {
	t.Run("should return list of orders without disscount", func(t *testing.T) {
		expected := []entity.Order{
			{ID: 1, TotalAmount: 47.5, CustomerID: 1},
			{ID: 2, TotalAmount: 95.0, CustomerID: 1},
			{ID: 3, TotalAmount: 950.0, CustomerID: 1},
		}

		orderRepo := mock.NewOrderRepositoryMock()
		orderRepo.ListOrdersFunc = func(filter input.OrderFilterInput) ([]entity.Order, error) {
			return []entity.Order{
				{ID: 1, TotalAmount: 50, CustomerID: 1},
				{ID: 2, TotalAmount: 100, CustomerID: 1},
				{ID: 3, TotalAmount: 1000, CustomerID: 1},
			}, nil
		}
		productRepo := mock.NewProductRepositoryMock()
		paymentMock := paymentmock.NewPaymentServiceMock()

		uc := usecase.NewOrderUseCase(orderRepo, productRepo, paymentMock)

		result, err := uc.ListOrders(input.OrderFilterInput{})
		
		assert.NoError(t, err)
		assert.NotEqual(t, expected, result)
	})

	t.Run("should return list of orders successfully applying disscount", func(t *testing.T) {
		expected := []entity.Order{
			{ID: 1, TotalAmount: 47.5, CustomerID: 1},
			{ID: 2, TotalAmount: 95.0, CustomerID: 1},
			{ID: 3, TotalAmount: 950.0, CustomerID: 1},
		}

		orderRepo := mock.NewOrderRepositoryMock()
		orderRepo.ListOrdersFunc = func(filter input.OrderFilterInput) ([]entity.Order, error) {
			return []entity.Order{
				{ID: 1, TotalAmount: 50, CustomerID: 1},
				{ID: 2, TotalAmount: 100, CustomerID: 1},
				{ID: 3, TotalAmount: 1000, CustomerID: 1},
			}, nil
		}
		productRepo := mock.NewProductRepositoryMock()
		paymentMock := paymentmock.NewPaymentServiceMock()

		uc := usecase.NewOrderUseCase(orderRepo, productRepo, paymentMock)

		id := 1
		result, err := uc.ListOrders(input.OrderFilterInput{CustomerID: &id})
		
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("should return list of orders successfully", func(t *testing.T) {
		expected := []entity.Order{
			{ID: 1, Status: "PAYMENT_PENDING", CustomerID: 10},
			{ID: 2, Status: "COMPLETED", CustomerID: 10},
		}

		orderRepo := mock.NewOrderRepositoryMock()
		orderRepo.ListOrdersFunc = func(filter input.OrderFilterInput) ([]entity.Order, error) {
			return expected, nil
		}
		productRepo := mock.NewProductRepositoryMock()
		paymentMock := paymentmock.NewPaymentServiceMock()

		uc := usecase.NewOrderUseCase(orderRepo, productRepo, paymentMock)

		result, err := uc.ListOrders(input.OrderFilterInput{Status: "PAYMENT_PENDING"})

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("should return error from repository", func(t *testing.T) {
		orderRepo := mock.NewOrderRepositoryMock()
		orderRepo.ListOrdersFunc = func(filter input.OrderFilterInput) ([]entity.Order, error) {
			return nil, assert.AnError
		}
		productRepo := mock.NewProductRepositoryMock()
		paymentMock := paymentmock.NewPaymentServiceMock()

		uc := usecase.NewOrderUseCase(orderRepo, productRepo, paymentMock)

		result, err := uc.ListOrders(input.OrderFilterInput{Status: "FAIL"})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestCheckout(t *testing.T) {
	t.Run("should update status to PAYMENT_RECEIVED when payment succeeds", func(t *testing.T) {
		orderRepo := mock.NewOrderRepositoryMock()
		orderRepo.UpdateStatusFunc = func(id int, status string) error {
			assert.Equal(t, "PAYMENT_RECEIVED", status)
			return nil
		}
		paymentMock := paymentmock.NewPaymentServiceMock()
		paymentMock.PaymentFunc = func(payment.PaymentInput) (payment.PaymentOutput, error) {
			return payment.PaymentOutput{}, nil
		}
		productRepo := mock.NewProductRepositoryMock()

		uc := usecase.NewOrderUseCase(orderRepo, productRepo, paymentMock)

		err := uc.Checkout(1)
		assert.NoError(t, err)
	})

	t.Run("should return error when payment fails", func(t *testing.T) {
		orderRepo := mock.NewOrderRepositoryMock()

		paymentMock := paymentmock.NewPaymentServiceMock()
		paymentMock.PaymentFunc = func(payment.PaymentInput) (payment.PaymentOutput, error) {
			return payment.PaymentOutput{}, errors.New("card declined")
		}

		productRepo := mock.NewProductRepositoryMock()

		uc := usecase.NewOrderUseCase(orderRepo, productRepo, paymentMock)

		err := uc.Checkout(1)
		assert.Error(t, err)
		assert.EqualError(t, err, "payment failed")
	})
}
