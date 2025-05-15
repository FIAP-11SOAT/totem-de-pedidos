package usecase

import (
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/domain/entity"
	"github.com/FIAP-11SOAT/totem-de-pedidos/internal/core/ports/inputs"
)

type Category interface {
	GetCategories() ([]*entity.Category, error)

	CreateCategory(categoryDTO *inputs.CategoryInput) (*entity.Category, error)

	UpdateCategory(categoryDTO *entity.Category) (*entity.Category, error)

	DeleteCategory(categoryID int) error

	FindCategoryById(categoryID int) (*entity.Category, error)
}
