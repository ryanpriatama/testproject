package service

import (
	"context"
	"database/sql"
	"ryan-test-backend/exception"
	"ryan-test-backend/helper"
	"ryan-test-backend/model/domain"
	"ryan-test-backend/model/web"
	"ryan-test-backend/repository"
	"time"

	"github.com/go-playground/validator"
)

//Implementation Business Logic or Service

type ProductServiceImpl struct {
	ProductRepostiroy repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepostiroy: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest, requestId int) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	timeNow := time.Now()

	product := domain.Product{
		Customer:  request.Customer,
		Quantity:  request.Quantity,
		Price:     request.Price,
		Timestamp: timeNow,
		RequestId: requestId,
	}

	product = service.ProductRepostiroy.Save(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId int, requestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.ProductRepostiroy.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepostiroy.Delete(ctx, tx, category)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int, requestId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.ProductRepostiroy.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(category)
}
