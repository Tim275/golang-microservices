package main

import (
	"log"
	"net"

	common "github.com/Tim275/golang-microservices/commons"
	"github.com/Tim275/golang-microservices/commons/api"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Laden der Umgebungsvariablen
	err := godotenv.Load()
	if err != nil {
		log.Println("Keine .env Datei gefunden")
	}

	grpcAddr := common.EnvString("GRPC_ADDR", ":2000")

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", grpcAddr, err)
	}

	grpcServer := grpc.NewServer()
	store := NewStore()
	svc := NewService(store)

	api.RegisterOrderServiceServer(grpcServer, svc)

	// Aktivieren Sie die gRPC-Reflection-API (nützlich für Debugging-Tools wie grpcurl)
	reflection.Register(grpcServer)

	log.Printf("Starting gRPC server on %s", grpcAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
