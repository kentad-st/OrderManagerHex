package model

type Order struct {
	Id *int64
	User *string
	Item *string
	Price *int64
	Quantity *int64
	OrderDate *string
}