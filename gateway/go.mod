module github.com/Tim275/golang-microservices/gateway

go 1.23.4

require (
	github.com/Tim275/golang-microservices/commons v0.0.0
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
)

// go.mod
replace github.com/Tim275/golang-microservices/commons => ../common
