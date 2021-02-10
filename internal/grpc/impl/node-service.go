package impl

import (
	"context"

	"github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"
	node "github.com/eddyjlhaigh/cardano-test-framework/internal/grpc/service"
)

type NodeServiceGrpcImpl struct {
}

func NewNodeServiceGrpcImpl() *NodeServiceGrpcImpl {
	return &NodeServiceGrpcImpl{}
}

func (serviceImpl *NodeServiceGrpcImpl) PrintDir(ctx context.Context, in *node.Request) (*service.Response, error) {
	return &node.Response{Body: "Test"}, nil
}

func (serviceImpl *NodeServiceGrpcImpl) FetchResponse(ctx context.Context, in *node.Request) (*service.Response, error) {
	return &node.Response{Body: "Test"}, nil
}
