package repositories

import (
	"context"

	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
)

type Category interface {
	CreateCategory(ctx context.Context, category entity.ProductCategory) (int, error)
}
