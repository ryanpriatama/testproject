package controller

import (
	"net/http"
	"ryan-test-backend/helper"
	"ryan-test-backend/model/web"
	"ryan-test-backend/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {
	requestId := request.URL.Query().Get("request_id")
	requestIdInt, _ := strconv.Atoi(requestId)

	productCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productCreateRequest)
	productResponse := controller.ProductService.Create(request.Context(), productCreateRequest, requestIdInt)

	webResponse := web.WebResponse{
		Code:      200,
		Status:    "OK",
		RequestId: requestIdInt,
		Data:      productResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := request.URL.Query().Get("id")
	id, _ := strconv.Atoi(productId)

	requestId := request.URL.Query().Get("request_id")
	requestIdInt, _ := strconv.Atoi(requestId)

	controller.ProductService.Delete(request.Context(), id, requestIdInt)
	webResponse := web.WebResponse{
		Code:      200,
		Status:    "OK",
		RequestId: requestIdInt,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := request.URL.Query().Get("id")
	id, _ := strconv.Atoi(productId)

	requestId := request.URL.Query().Get("request_id")
	requestIdInt, _ := strconv.Atoi(requestId)

	categoryResponse := controller.ProductService.FindById(request.Context(), id, requestIdInt)
	webResponse := web.WebResponse{
		Code:      200,
		Status:    "OK",
		RequestId: requestIdInt,
		Data:      categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
