package service

import (
	"database/sql"
	"reflect"
	"ryan-test-backend/repository"
	"testing"

	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	//"gopkg.in/go-playground/assert.v1"
)

func TestNewProductService(t *testing.T) {
	type args struct {
		productRepository repository.ProductRepository
		DB                *sql.DB
		validate          *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want ProductService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductService(tt.args.productRepository, tt.args.DB, tt.args.validate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductService() = %v, want %v", got, tt.want)
			}
		})
	}
}
