package helper

import (
	"ryan-test-backend/model/domain"
	"testing"
)

func TestToProductResponse(t *testing.T) {
	// Arrange
	product := domain.Product{
		Id:       1,
		Customer: "Product 1",
		Price:    1000,
		Quantity: 10,
	}

	// Act
	response := ToProductResponse(product)

	// Assert
	if response.Id != product.Id {
		t.Errorf("Incorrect response id, got: %d, want: %d", response.Id, product.Id)
	}
	if response.Customer != product.Customer {
		t.Errorf("Incorrect response name, got: %s, want: %s", response.Customer, product.Customer)
	}
	if response.Price != product.Price {
		t.Errorf("Incorrect response price, got: %.2f, want: %.2f", response.Price, product.Price)
	}
	if response.Quantity != product.Quantity {
		t.Errorf("Incorrect response quantity, got: %d, want: %d", response.Quantity, product.Quantity)
	}
}

func TestToProductResponses(t *testing.T) {
	// Arrange
	products := []domain.Product{
		{
			Id:       1,
			Customer: "Product 1",
			Price:    1000,
			Quantity: 10,
		},
		{
			Id:       2,
			Customer: "Product 2",
			Price:    2000,
			Quantity: 20,
		},
	}

	// Act
	responses := ToProductResponses(products)

	// Assert
	if len(responses) != len(products) {
		t.Errorf("Incorrect number of responses, got: %d, want: %d", len(responses), len(products))
	}

	for i, response := range responses {
		if response.Id != products[i].Id {
			t.Errorf("Incorrect response id, got: %d, want: %d", response.Id, products[i].Id)
		}
		if response.Customer != products[i].Customer {
			t.Errorf("Incorrect response name, got: %s, want: %s", response.Customer, products[i].Customer)
		}
		if response.Price != products[i].Price {
			t.Errorf("Incorrect response price, got: %.2f, want: %.2f", response.Price, products[i].Price)
		}
		if response.Quantity != products[i].Quantity {
			t.Errorf("Incorrect response quantity, got: %d, want: %d", response.Quantity, products[i].Quantity)
		}
	}
}
