package main

import (
	"log"
	"net/http"

	common "github.com/Tim275/golang-microservices/commons"
	"github.com/Tim275/golang-microservices/commons/api"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Laden der Umgebungsvariablen
	err := godotenv.Load()
	if err != nil {
		log.Println("Keine .env Datei gefunden")
	}

	httpAddr := common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr := common.EnvString("ORDER_SERVICE_ADDR", "localhost:2000")

	// Verbindung zum gRPC OrderService herstellen
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	defer conn.Close()

	orderClient := api.NewOrderServiceClient(conn)

	// HTTP Router einrichten
	router := mux.NewRouter()
	handler := NewHandler(orderClient)
	handler.registerRoutes(router)

	// HTTP Server starten
	log.Printf("Starting HTTP server on %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, router); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
