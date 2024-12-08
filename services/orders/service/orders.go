package service

import (
	"context"

	"github.com/Bossnicks/go-grpc-microservice-demo/services/common/genproto/orders"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	//store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(context.Context, *orders.Order) error {
	ordersDb = append(ordersDb, order)
	return nil
}
