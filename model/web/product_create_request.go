package web

type ProductCreateRequest struct {
	Customer  string  `validate:"required,max=300,min=1" json:"customer"`
	Price     float32 `validate:"required,min=1" json:"price"`
	Quantity  int     `validate:"required" json:"quantity"`
	RequestId int     `json:"request_id"`
}
