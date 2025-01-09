// orders/service.go
package main

import (
	"context"

	"github.com/Tim275/golang-microservices/commons/api"
)

// service implementiert den OrderServiceServer
type service struct {
	api.UnimplementedOrderServiceServer
	store OrderStore
}

// NewService erstellt einen neuen Service mit dem gegebenen OrderStore
func NewService(store OrderStore) *service {
	return &service{store: store}
}

// CreateOrder verarbeitet die Erstellung einer Bestellung
func (s *service) CreateOrder(ctx context.Context, req *api.CreateOrderRequest) (*api.CreateOrderResponse, error) {
	order := &api.Order{
		ID:         generateOrderID(), // Implementieren Sie eine ID-Generierung
		CustomerID: req.CustomerID,
		Status:     "created",
		Items:      make([]*api.OrderItem, 0),
	}

	for _, item := range req.Items {
		orderItem := &api.OrderItem{
			ID:       item.ItemID,
			Quantity: item.Quantity,
			PriceID:  "", // Setzen Sie dies nach Bedarf
		}
		order.Items = append(order.Items, orderItem)
	}

	// Bestellung im Store speichern
	err := s.store.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	resp := &api.CreateOrderResponse{
		Status:       "success",
		Message:      "Order created successfully",
		CreatedOrder: order,
	}

	return resp, nil
}

// GenerateOrderID generiert eine eindeutige Order-ID
func generateOrderID() string {
	// Implementieren Sie eine sinnvolle ID-Generierung, z.B. UUID
	return "1" // Für Einfachheit
}
