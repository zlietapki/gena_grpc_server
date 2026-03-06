package grpc_handler

import (
	"context"

	"github.com/zlietapki/microboiler_api_contracts/pkg/pb/v1"
)

func (h *Handler) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRes, error) {
	return &pb.HelloRes{}, nil
}
