// main.go
package main

import (
	"log"
	"net/http"

	common "github.com/Tim275/golang-microservices/commons"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Keine .env Datei gefunden")
	}

	httpAddr := common.EnvString("HTTP_ADDR", ":8080")

	router := mux.NewRouter()
	handler := NewHandler()
	handler.registerRoutes(router)

	log.Printf("Starting server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, router); err != nil {
		log.Fatal("Failed to start http server: ", err)
	}
}
