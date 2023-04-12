package repository

import (
	"context"
	"database/sql"
	"errors"
	"ryan-test-backend/helper"
	"ryan-test-backend/model/domain"
	"strconv"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

//Implementation from contract interface ProductRepository

func (f *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	var id string
	SQL := "INSERT INTO new_products (customer,quantity,price,timestamp,request_id) VALUES($1,$2,$3,$4,$5) RETURNING id"
	err := tx.QueryRowContext(ctx, SQL, product.Customer, product.Quantity, product.Price, product.Timestamp, product.RequestId).Scan(&id)
	helper.PanicIfError(err)

	product.Id, _ = strconv.Atoi(id)
	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "DELETE FROM new_products WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "select id, customer, quantity, price, timestamp, request_id from new_products where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Customer, &product.Quantity, &product.Price, &product.Timestamp, &product.RequestId)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("category is not found")
	}
}
