package mocks

import (
	"context"
	"ryan-test-backend/model/domain"
	"ryan-test-backend/model/web"
	"testing"

	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/assert.v1"
)

func TestMockProductService_Create(t *testing.T) {
	mockProduct := domain.Product{
		Id:       1,
		Customer: "Product 1",
		Price:    1000,
		Quantity: 10,
	}
	mockProductService := new(MockProductService)
	mockProductService.On("Create", mock.Anything, mock.Anything).Return(mockProduct)

	ctx := context.Background()
	request := web.ProductCreateRequest{
		Customer: "Product 1",
		Price:    1000,
		Quantity: 10,
	}
	product := mockProductService.Create(ctx, request)

	assert.Equal(t, mockProduct, product)
	mockProductService.AssertExpectations(t)
}

func TestMockProductRepository_Save(t *testing.T) {
	repo := &MockProductRepository{}
	product := domain.Product{Customer: "Product D", Price: 4000, Quantity: 40}
	savedProduct := repo.Save(context.Background(), nil, product)

	if savedProduct != product {
		t.Errorf("Save() = %v, want %v", savedProduct, product)
	}
}
