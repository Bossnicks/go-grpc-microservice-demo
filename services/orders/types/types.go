package types

import "context"

type OrderService interface {
	CreateOrder(ctx context.Context, *orders.Order) error
}
