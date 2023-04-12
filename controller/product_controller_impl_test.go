package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"ryan-test-backend/app"
	"ryan-test-backend/mocks"
	"ryan-test-backend/model/web"
	"ryan-test-backend/repository"
	"ryan-test-backend/service"
	"testing"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/assert.v1"
)

func TestProductController_Create(t *testing.T) {
	// Prepare mock ProductService
	mockProductService := &mocks.MockProductService{}
	mockProduct := web.ProductResponse{
		Id:       22,
		Customer: "test",
		Price:    100,
		Quantity: 10,
	}
	mockProductService.On("Create", mock.Anything, mock.Anything).Return(mockProduct)

	// Prepare request body
	requestBody := web.ProductCreateRequest{
		Customer: "test",
		Price:    100,
		Quantity: 10,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	// Prepare request
	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	db := app.NewDB()
	//test.TruncateCategory(db)
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)

	// Prepare controller and router
	controller := &ProductControllerImpl{
		ProductService: productService,
	}
	router := httprouter.New()
	router.POST("/products", controller.Create)

	// Make request and check response
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var response web.WebResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	respDataMap := response.Data.(map[string]interface{})

	ResponseData := web.ProductResponse{
		Id:       int(respDataMap["id"].(float64)),
		Customer: respDataMap["customer"].(string),
		Price:    float32(respDataMap["price"].(float64)),
		Quantity: int(respDataMap["quantity"].(float64)),
	}
	assert.Equal(t, mockProduct, ResponseData)
}

func TestNewProductController(t *testing.T) {
	db := app.NewDB()
	//test.TruncateCategory(db)
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := NewProductController(productService)

	type args struct {
		productService service.ProductService
	}
	tests := []struct {
		name string
		args args
		want ProductController
	}{
		{
			name: "default",
			args: args{
				productService: productService,
			},
			want: productController,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductController(tt.args.productService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductController() = %v, want %v", got, tt.want)
			}
		})
	}
}
