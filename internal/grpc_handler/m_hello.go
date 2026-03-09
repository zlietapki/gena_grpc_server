package grpc_handler

import (
	"context"

	"github.com/zlietapki/microboiler_api_contracts/pkg/pb/v1"
)

func (h *Handler) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRes, error) {
	resp, err := h.uc.Hello(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &pb.HelloRes{
		Hello: resp,
	}, nil
}
