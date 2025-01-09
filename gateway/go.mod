module github.com/Tim275/golang-microservices/gateway

go 1.23.4

require (
	github.com/Tim275/golang-microservices/commons v0.0.0
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	google.golang.org/grpc v1.69.2
	google.golang.org/protobuf v1.35.1 // indirect
)

require (
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
)

replace github.com/Tim275/golang-microservices/commons => ../common
