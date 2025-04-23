package repositories

import (
	"errors"
	"sync"
	"time"
	"totem-pedidos/internal/core/domain/entity"
)

type MemoryCustomerRepository struct {
	customers map[string]*entity.Customer
	mu        sync.RWMutex
	nextID    int
}

func NewMemoryCustomerRepository() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[string]*entity.Customer),
		nextID:    1,
	}
}

func (r *MemoryCustomerRepository) GetByTaxID(taxID string) (*entity.Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	customer, exists := r.customers[taxID]
	if !exists {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

func (r *MemoryCustomerRepository) Create(customer *entity.Customer) (*entity.Customer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verificar se já existe um cliente com o mesmo TaxID
	for _, existingCustomer := range r.customers {
		if existingCustomer.TaxID == customer.TaxID {
			return nil, errors.New("customer with this tax ID already exists")
		}
	}

	// Atribuir ID e timestamp
	customer.ID = r.nextID
	customer.CreatedAt = time.Now()
	r.nextID++

	// Armazenar o cliente
	r.customers[customer.TaxID] = customer

	return customer, nil
}

func (r *MemoryCustomerRepository) Update(customer *entity.Customer) (*entity.Customer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, exists := r.customers[customer.TaxID]
	if !exists {
		return nil, errors.New("customer not found")
	}

	// Preservar ID e CreatedAt originais
	customer.ID = existing.ID
	customer.CreatedAt = existing.CreatedAt

	// Atualizar o cliente
	r.customers[customer.TaxID] = customer

	return customer, nil
}
