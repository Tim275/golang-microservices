package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(router *mux.Router) {
	// Statische Dateien dienen
	fileServer := http.FileServer(http.Dir("public"))
	router.PathPrefix("/").Handler(fileServer)

	// API-Routen
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/customers/{customerID}/orders", h.HandleCreateOrder).Methods("POST")
	api.HandleFunc("/customers/{customerID}/orders/{orderID}", h.HandleGetOrder).Methods("GET")
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// Implementiere die Logik zum Erstellen einer Bestellung
}

func (h *handler) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	// Implementiere die Logik zum Abrufen einer Bestellung
}
