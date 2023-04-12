package service

import (
	"context"
	"ryan-test-backend/model/web"
)

//contract interface business logic / service

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest, requestId int) web.ProductResponse
	Delete(ctx context.Context, productId int, requestId int)
	FindById(ctx context.Context, productId int, requestId int) web.ProductResponse
}
