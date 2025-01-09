// orders/store.go
package main

import (
	"context"
	"sync"

	"github.com/Tim275/golang-microservices/commons/api"
)

// store implementiert OrderStore
type store struct {
	mu     sync.Mutex
	orders map[string]*api.Order
}

// NewStore erstellt einen neuen Store
func NewStore() OrderStore {
	return &store{
		orders: make(map[string]*api.Order),
	}
}

// Create fügt eine neue Bestellung hinzu (einfaches In-Memory-Store)
func (s *store) Create(ctx context.Context, order *api.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.orders[order.ID] = order
	return nil
}

// Optional: Implementieren Sie weitere Methoden wie Get, Update, Delete
