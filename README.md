# Cardano Test Framework Prototype

A prototype to see if gRPC would be a valid tool to help with Cardano test frameworks.

## Requirements:

- Golang 1.15
- Protobuf Compiler 3.0

Installation notes can be found here https://grpc.io/docs/languages/go/quickstart/

## Usage

### Building and Running

`./proto-gen.sh` compiles `.proto` files
`./run-server.sh <port>` and `./run-client.sh <port>` runs a server and client.
`./build-server.sh <port>` and `./build-client.sh <port>` builds a server and client.

### Contributing

There are five key files. 

- `./internal/proto-files/service/node-service.proto`
- `./internal/grpc/service/node-service.pb.go`
- `./internal/grpc/impl/node-service.go`
- `./cmd/grpc/server/main.go`
- `./cmd/grpc/client/main.go`

`node-service.proto` contains the data structures and gRPC methods that you want to have. By compiling these with `./proto-gen.sh`, `node-service.pb.go` is generated and used by both the client and server. Note then that the `node-service.proto` is coupled between the two services. On the server, `node-service.go` implements the methods in `node-service.proto` and this is then registered in `server/main.go`. `client/main.go` has a copy of `node-service.pb.go` so that it knows what methods are available and what data types they require/respond wtih, but no other details are expected. 

Therefore to extend this project with a new gRPC method, you would need to add it to `node-service.proto` along with its request and response arguments types (a string is acceptable!). Compiling this with `./proto-gen.sh` will create a Golang file however you can install a different proto compiler to generate other languages. Following this, you would need to implement the function in `node-service.go` as per the other methods already there. Finally, you would want to add this function to `client/main.go` so that you can use it. 
