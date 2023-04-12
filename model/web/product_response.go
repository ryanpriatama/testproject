package web

import "time"

type ProductResponse struct {
	Id        int       `json:"id"`
	Customer  string    `json:"customer"`
	Price     float32   `json:"price"`
	Quantity  int       `json:"quantity"`
	Timestamp time.Time `json:"timestamp"`
}
