package service

import (
	model "OrderManagerHex/domain/model"

	repository "OrderManagerHex/domain/repository"

	"time"
)

type OrderManagementService interface {
	ShowOrders(string, string) ([]model.Order, error)
	AddOrder(*[]model.Order)
}

type orderManagementService struct {
	orderRepository repository.OrderRepository
}

func NewOrderManagementService(cr repository.OrderRepository) OrderManagementService {
	return &orderManagementService{
		orderRepository :cr,
	}
}

func (cs orderManagementService) ShowOrders(user string, orderDate string) (items []model.Order, err error) {
	items, err = cs.orderRepository.GetOrders(user, orderDate)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func(cs orderManagementService) AddOrder(orders *[]model.Order) {
    day := time.Now()
	const layout = "2006-01-02"
	daystr := day.Format(layout)
	for _, v := range *orders {
		v.OrderDate = &daystr
		cs.orderRepository.Post(&v)
	}
}