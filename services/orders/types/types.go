package types

import (
	"context"

	"github.com/Bossnicks/go-grpc-microservice-demo/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
