# Learning gRPC

# Compile Protos
```
// Generate protobuf structs
protoc --go_opt=module=github.com/Saddham/learning-grpc  --go_out=. ./rides/pb/rides.proto

// Generate protobuf structs and grpc service specification
protoc --go_opt=module=github.com/Saddham/learning-grpc  --go_out=. --go-grpc_out=. --go-grpc_opt=module=github.com/Saddham/learning-grpc  ./rides/pb/rides.proto

// Generate protobuf structs, grpc service specification and grpc gateway
protoc --go_opt=module=github.com/Saddham/learning-grpc  --go_out=. --go-grpc_out=. --go-grpc_opt=module=github.com/Saddham/learning-grpc --grpc-gateway_out=. --grpc-gateway_opt=module=github.com/Saddham/learning-grpc --grpc-gateway_opt=generate_unbound_methods=true ./rides/pb/rides.proto
```

# Go Commands Used
```
go mod init
go mod download google.golang.org/grpc
go mod download google.golang.org/protobuf
go mod tidy
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
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

// Run grpc gateway
go run rides/gateway/main.go

// Call gateway http api
curl http://localhost:8080/Rides/Start -d@rides/data/start.json
```

# Test Server
```
go test learning-grpc/rides/server
go test -run ^TestEnd$ learning-grpc/rides/server
go test -run ^TestEndE2E$ learning-grpc/rides/server
```