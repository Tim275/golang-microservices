// orders/types.go
package main

import (
	"context"

	"github.com/Tim275/golang-microservices/commons/api"
)

type OrderStore interface {
	Create(ctx context.Context, order *api.Order) error
	// Weitere Methoden wie Get, Update, Delete können hinzugefügt werden
}

// Sie können die OrderService-Schnittstelle entfernen, da Sie den gRPC-Server bereits implementieren
