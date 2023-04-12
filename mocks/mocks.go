package mocks

import (
	"context"
	"database/sql"
	"ryan-test-backend/model/domain"
	"ryan-test-backend/model/web"

	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) Create(ctx context.Context, request web.ProductCreateRequest) domain.Product {
	args := m.Called(ctx, request)
	return args.Get(0).(domain.Product)
}

type MockProductRepository struct{}

func (m *MockProductRepository) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	return product
}

func (m *MockProductRepository) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
}

func (m *MockProductRepository) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	var product domain.Product
	return product, nil
}
