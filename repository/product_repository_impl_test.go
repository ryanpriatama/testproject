package repository

import (
	"context"
	"database/sql"
	"log"
	"ryan-test-backend/model/domain"
	"ryan-test-backend/test"
	"testing"

	_ "github.com/lib/pq"
)

func TestProductRepositoryImpl_Save(t *testing.T) {
	repo := NewProductRepository()

	// setup database connection
	connStr := "postgres://postgres:rootadmin@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	test.TruncateCategory(db)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// create transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	// create product to be saved
	product := domain.Product{
		Customer: "Test Product",
		Price:    1000,
		Quantity: 10,
	}

	// save product
	ctx := context.Background()
	savedProduct := repo.Save(ctx, tx, product)

	// check if product was saved correctly
	if savedProduct.Id == 0 {
		t.Errorf("Expected saved product ID to be non-zero")
	}
}
