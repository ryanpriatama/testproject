package domain

import "time"

//Domain or entity is representation field on dabase

type Product struct {
	Id        int
	Customer  string
	Quantity  int
	Price     float32
	Timestamp time.Time
	RequestId int
}
