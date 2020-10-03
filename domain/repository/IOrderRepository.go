package repository

import (

	model "OrderManagerHex/domain/model"
)

type OrderRepository interface {
	Post(*model.Order)
	GetOrders(user string, orderDate string) ([]model.Order, error)
}