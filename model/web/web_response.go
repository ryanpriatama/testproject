package web

type WebResponse struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	RequestId int         `json:"request_id"`
	Data      interface{} `json:"data"`
}
