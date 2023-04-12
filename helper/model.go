package helper

import (
	"ryan-test-backend/model/domain"
	"ryan-test-backend/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:        product.Id,
		Customer:  product.Customer,
		Price:     product.Price,
		Quantity:  product.Quantity,
		Timestamp: product.Timestamp,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
