package main

import (
	handler "OrderManagerHex/adapter/http"
	postgres "OrderManagerHex/adapter/postgres"
	service "OrderManagerHex/service"
)

func main(){
	orderPersistence := postgres.NewOrderPersistence()
	orderService := service.NewOrderManagementService(orderPersistence)
	orderHandler := handler.NewOrderHandler(orderService)

	orderHandler.MainHttpd()
}