package repositories

import (
	"testing"
	"totem-pedidos/internal/core/domain/entity"
)

func TestMemoryCustomerRepository_GetByTaxId(t *testing.T) {
	repo := NewMemoryCustomerRepository()
	customer := &entity.Customer{
		Name:  "John Doe",
		Email: "john@example.com",
		TaxID: "123456789",
	}

	// First create a customer to test with
	createdCustomer, err := repo.Create(customer)
	if err != nil {
		t.Fatalf("Failed to create test customer: %v", err)
	}

	tests := []struct {
		name    string
		taxID   string
		want    *entity.Customer
		wantErr bool
	}{
		{
			name:    "Existing customer",
			taxID:   "123456789",
			want:    createdCustomer,
			wantErr: false,
		},
		{
			name:    "Non-existing customer",
			taxID:   "999999999",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetByTaxID(tt.taxID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByTaxId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.TaxID != tt.want.TaxID {
				t.Errorf("GetByTaxId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryCustomerRepository_Create(t *testing.T) {
	repo := NewMemoryCustomerRepository()

	tests := []struct {
		name     string
		customer *entity.Customer
		wantErr  bool
	}{
		{
			name: "Create new customer",
			customer: &entity.Customer{
				Name:  "John Doe",
				Email: "john@example.com",
				TaxID: "123456789",
			},
			wantErr: false,
		},
		{
			name: "Create duplicate customer",
			customer: &entity.Customer{
				Name:  "Jane Doe",
				Email: "jane@example.com",
				TaxID: "123456789", // Same TaxID as above
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.Create(tt.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.ID == 0 {
					t.Error("Create() ID not set")
				}
				if got.CreatedAt.IsZero() {
					t.Error("Create() CreatedAt not set")
				}
				if got.TaxID != tt.customer.TaxID {
					t.Errorf("Create() TaxID = %v, want %v", got.TaxID, tt.customer.TaxID)
				}
			}
		})
	}
}

func TestMemoryCustomerRepository_Update(t *testing.T) {
	repo := NewMemoryCustomerRepository()

	// Create initial customer
	original := &entity.Customer{
		Name:  "John Doe",
		Email: "john@example.com",
		TaxID: "123456789",
	}
	created, err := repo.Create(original)
	if err != nil {
		t.Fatalf("Failed to create test customer: %v", err)
	}

	tests := []struct {
		name     string
		customer *entity.Customer
		wantErr  bool
	}{
		{
			name: "Update existing customer",
			customer: &entity.Customer{
				ID:    created.ID,
				Name:  "John Updated",
				Email: "john.updated@example.com",
				TaxID: created.TaxID,
			},
			wantErr: false,
		},
		{
			name: "Update non-existing customer",
			customer: &entity.Customer{
				Name:  "Non Existent",
				Email: "non@example.com",
				TaxID: "999999999",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.Update(tt.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Name != tt.customer.Name {
					t.Errorf("Update() Name = %v, want %v", got.Name, tt.customer.Name)
				}
				if got.Email != tt.customer.Email {
					t.Errorf("Update() Email = %v, want %v", got.Email, tt.customer.Email)
				}
				if got.CreatedAt != created.CreatedAt {
					t.Error("Update() CreatedAt was modified")
				}
			}
		})
	}
}
