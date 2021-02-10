protoc --go_out=plugins=grpc:./ ./internal/proto-files/service/node-service.proto
mv ./internal/proto-files/service/node-service.pb.go ./internal/grpc/service