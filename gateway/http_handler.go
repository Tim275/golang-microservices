package main

import "net/http"

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// Implementiere die Logik zum Erstellen einer Bestellung
}

func (h *handler) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	// Implementiere die Logik zum Abrufen einer Bestellung
}
