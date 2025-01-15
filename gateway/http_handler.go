package main

import (
	"log"
	"net/http"

	common "github.com/Tim275/golang-microservices/commons"
	pb "github.com/Tim275/golang-microservices/commons/api"
	"github.com/gorilla/mux"
)

// handler strukturiert den OrderServiceClient
type handler struct {
	orderClient pb.OrderServiceClient
}

// NewHandler erstellt einen neuen Handler mit dem gegebenen OrderServiceClient
func NewHandler(orderClient pb.OrderServiceClient) *handler {
	return &handler{orderClient: orderClient}
}

// registerRoutes registriert die API-Routen zuerst, dann die statischen Dateien
func (h *handler) registerRoutes(router *mux.Router) {
	// API-Routen zuerst registrieren
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/customers/{customerID}/orders", h.HandleCreateOrder).Methods("POST")
	apiRouter.HandleFunc("/customers/{customerID}/orders/{orderID}", h.HandleGetOrder).Methods("GET")

	// Statische Dateien bedienen
	fileServer := http.FileServer(http.Dir("public"))
	router.PathPrefix("/").Handler(fileServer)
}

// HandleCreateOrder verarbeitet die Erstellung einer Bestellung
// ...existing code...

// HandleCreateOrder verarbeitet die Erstellung einer Bestellung
func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customerID"]

	var items []*pb.ItemsWithQuantity
	if err := common.ReadJson(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.orderClient.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	if err != nil {
		log.Printf("Failed to create order: %v", err)
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	common.WriteJSON(w, http.StatusOK, resp)
}

// HandleGetOrder verarbeitet das Abrufen einer Bestellung (optional)
func (h *handler) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderID"]

	resp, err := h.orderClient.GetOrder(r.Context(), &pb.GetOrderRequest{
		OrderID: orderID,
	})
	if err != nil {
		log.Printf("Failed to get order: %v", err)
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	common.WriteJSON(w, http.StatusOK, resp)
}

// ...existing code...
