# Learning gRPC

# Compile Protos
```
// Generate protobuf structs
protoc --go_opt=module=github.com/Saddham/learning-grpc  --go_out=. ./rides/pb/rides.proto

// Generate protobuf structs and grpc service specification
protoc --go_opt=module=github.com/Saddham/learning-grpc  --go_out=. --go-grpc_out=. --go-grpc_opt=module=github.com/Saddham/learning-grpc  ./rides/pb/rides.proto
```

# Go Commands Used
```
go mod init
go mod download google.golang.org/grpc
go mod download google.golang.org/protobuf
go mod tidy
```

# GRPC Server & Client
```
// List Services
grpcurl --plaintext localhost:9292 list

// List Service APIs
grpcurl --plaintext localhost:9292 list Rides

// Describe Service APIs
grpcurl --plaintext localhost:9292 describe Rides
grpcurl --plaintext localhost:9292 describe .StartRequest

// Call Service API
grpcurl --plaintext -d @ localhost:9292 Rides.Start < rides/data/start.json

// Run grpc server
go run rides/server/main.go

// Run grpc client
go run rides/client/main.go 
```