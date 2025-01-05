// types.go
package main

import "context"

type OrderStore interface {
	Create(ctx context.Context) error
}

type OrderService interface {
	CreateOrder(ctx context.Context) error
}
